package membership

import (
	"context"
	"sync"
	"time"

	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/memberlist"
	"github.com/sakari-ai/moirai/log"
	"github.com/sakari-ai/moirai/log/field"
)

type ClusterMember struct {
	Address  string
	master   string
	election *Election

	connected bool
	mux       *sync.Mutex

	msgCh      chan []byte
	broadcasts *memberlist.TransmitLimitedQueue
	cluster    *memberlist.Memberlist
}

func NewClusterMember(address string) *ClusterMember {
	return &ClusterMember{Address: address, mux: new(sync.Mutex), msgCh: make(chan []byte)}
}

//Membership Delegate
func (n *ClusterMember) NodeMeta(limit int) []byte {
	// not use, noop
	return []byte("")
}

//Membership Delegate
func (n *ClusterMember) NotifyMsg(msg []byte) {
	n.msgCh <- msg
}

//Membership Delegate
func (n *ClusterMember) GetBroadcasts(overhead, limit int) [][]byte {
	return n.broadcasts.GetBroadcasts(overhead, limit)
}

//Membership Delegate
func (n *ClusterMember) LocalState(join bool) []byte {
	// not use, noop
	return []byte("")
}

//Membership Delegate
func (n *ClusterMember) MergeRemoteState(buf []byte, join bool) {
	// not use
}

func (n *ClusterMember) EventLeader(isMaster bool, master string) {
	if n.connected {
		return
	}
	if isMaster {
		n.master = master
		err := n.masterNode()
		if err != nil {
			log.Error("can not process master", field.Error(err))
		}
		return
	}
	slaver := false
	if master != "" && master != n.master && !isMaster {
		n.master = master
		slaver = true
	}
	if slaver {
		err := n.slaveNode()
		if err != nil {
			log.Error("can not process slaver", field.Error(err))
		}
	}
}

func (n *ClusterMember) Leave(node *memberlist.Node) {
	if node.String() == n.Address {
		n.connected = false
		attempt := 0
		for {
			err := n.slaveNode()
			attempt++
			if err == nil {
				break
			}
			if attempt > 10 {
				panic(err)
			}
		}
	}
}

func (n *ClusterMember) Broadcast(ctx context.Context, listener func([]byte)) {
	conf := api.DefaultConfig()
	conf.Address = "consul-server:8500"
	consul, _ := api.NewClient(conf)

	elconf := &ElectionConfig{
		CheckTimeout: 15 * time.Second,
		Client:       consul,
		Checks:       []string{},
		Key:          "service/volante/leader",
		Event:        n,
		Node:         n.Address,
	}
	e := NewElection(elconf)

	n.election = e
	go e.Init()

	for {
		select {
		case msg, ok := <-n.msgCh:
			if !ok {
				return
			}
			listener(msg)
		}
	}
}

func (n *ClusterMember) Emit(value []byte) error {
	for _, node := range n.cluster.Members() {
		if node.Name == n.Address {
			log.Info("Ignore member", field.String("cluster member Name", node.Name), field.String("node Address", n.Address))
			continue // skip self
		}
		_ = n.cluster.SendReliable(node, value)
	}
	return nil
}

func (n *ClusterMember) createNode() *memberlist.Memberlist {
	if n.cluster == nil {
		conf := memberlist.DefaultLocalConfig()
		conf.Name = n.Address
		conf.BindPort = 7947 // avoid port conflict
		conf.AdvertisePort = conf.BindPort
		conf.Events = n.election
		conf.Delegate = n
		n.broadcasts = new(memberlist.TransmitLimitedQueue)
		n.broadcasts.NumNodes = func() int {
			log.Info("broadcast nodes", field.Int("no#", n.election.Num))
			return n.election.Num
		}

		list, _ := memberlist.Create(conf)

		n.cluster = list
	}
	return n.cluster
}

func (n *ClusterMember) masterNode() error {
	n.mux.Lock()
	defer n.mux.Unlock()
	list := n.createNode()

	local := list.LocalNode()
	log.Info("Master node", field.String("ip", local.Addr.To4().String()), field.Int32("port", int32(local.Port)))

	n.connected = true
	return nil
}

func (n *ClusterMember) slaveNode() error {
	n.mux.Lock()
	defer n.mux.Unlock()
	list := n.createNode()

	if _, err := list.Join([]string{n.master}); err != nil {
		return err
	}
	n.connected = true
	return nil
}
