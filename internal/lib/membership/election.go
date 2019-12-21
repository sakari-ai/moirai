package membership

import (
	"encoding/base64"
	"runtime"
	"sync"
	"time"

	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/memberlist"
	"github.com/sakari-ai/moirai/log"
	"github.com/sakari-ai/moirai/log/field"
)

// Election implements to detect a leader in a cluster of services
type Election struct {
	Client       *api.Client // Consul client
	Checks       []string    // Slice of associated health checks
	leader       bool        // Flag of a leader
	Kv           string      // Key in Consul kv
	sessionID    string      // Id of session
	logLevel     uint8       //  Log level LogDisable|LogError|LogInfo|LogDebug
	inited       bool        // Flag of init.
	CheckTimeout time.Duration
	LogPrefix    string        // Prefix for a log
	stop         chan struct{} // chnnel to stop process
	success      chan struct{} // channel for the signal that the process is stopped
	Event        Notifier
	Node         string
	Master       string
	sync.RWMutex

	Num int
}

// Notifier can tell your code the event of the leader's status change
type Notifier interface {
	EventLeader(e bool, master string) // The method will be called when the leader status is changed
	Leave(node *memberlist.Node)
}

// ElectionConfig config for Election
type ElectionConfig struct {
	Client       *api.Client // Consul client
	Checks       []string    // Slice of associated health checks
	Key          string      // Key in Consul KV
	LogLevel     uint8       // Log level LogDisable|LogError|LogInfo|LogDebug
	LogPrefix    string      // Prefix for a log
	Event        Notifier
	CheckTimeout time.Duration
	Node         string
}

// IsLeader check a leader
func (e *Election) IsLeader() bool {
	e.RLock()
	defer e.RUnlock()
	return e.leader
}

func (e *Election) NotifyJoin(node *memberlist.Node) {
	e.Num += 1
}

func (e *Election) NotifyLeave(node *memberlist.Node) {
	e.Num -= 1
	log.Info("member_list_leave", field.Any("node", node))
	e.Event.Leave(node)
}

func (e *Election) NotifyUpdate(node *memberlist.Node) {
}

// SetLogLevel is setting level according constants LogDisable|LogError|LogInfo|LogDebug
func (e *Election) SetLogLevel(level uint8) {
	e.logLevel = level
}

// Params: Consul client, slice of associated health checks, service name
func NewElection(c *ElectionConfig) *Election {
	e := &Election{
		Client:       c.Client,
		Checks:       append(c.Checks, "serfHealth"),
		leader:       false,
		Kv:           c.Key,
		CheckTimeout: c.CheckTimeout,
		LogPrefix:    c.LogPrefix,
		stop:         make(chan struct{}),
		success:      make(chan struct{}),
		Node:         c.Node,
		Event:        c.Event,
	}
	return e
}

func (e *Election) createSession() (err error) {
	ses := &api.SessionEntry{
		Checks:    e.Checks,
		LockDelay: 15 * time.Second,
		TTL:       "20s",
		Name:      e.Node,
	}
	e.sessionID, _, err = e.Client.Session().Create(ses, nil)
	if err != nil {
		e.logError("Create session error " + err.Error())
	}
	return
}

func (e *Election) checkSession() (bool, error) {
	if e.sessionID == "" {
		return false, nil
	}
	res, _, err := e.Client.Session().Info(e.sessionID, nil)

	if err != nil {
		e.logError("Info session error " + err.Error())
	}

	return res != nil, err
}

func (e *Election) getMasterNode(node string) (string, error) {
	res, _, err := e.Client.Session().Info(node, nil)

	if res != nil {
		return res.Name, nil
	}
	return "", err
}

// Try to acquire
func (e *Election) acquire() (bool, error) {
	kv := &api.KVPair{
		Key:     e.Kv,
		Session: e.sessionID,
		Value:   []byte(e.sessionID),
	}
	res, _, err := e.Client.KV().Acquire(kv, nil)
	if err != nil {
		e.logError("Acquire kv error " + err.Error())
	}
	return res, err
}

func (e *Election) disableLeader() {
	e.Lock()
	defer e.Unlock()
	if e.leader {
		e.leader = false
	}
	if e.Event != nil {
		e.Event.EventLeader(e.leader, e.Master)
	}
}

func (e *Election) getKvSession() (string, error) {
	p, _, err := e.Client.KV().Get(e.Kv, nil)
	if err != nil {
		e.logError("Kv error " + err.Error())
		return "", err
	}
	if p == nil {
		return "", nil
	}
	encodedBase64 := p.Value
	s := base64.StdEncoding.EncodeToString(encodedBase64)

	// decode string to bytes
	b, err := base64.StdEncoding.DecodeString(s)
	master, err := e.getMasterNode(string(b))
	e.Master = master
	return p.Session, err
}

// Init starting election process
func (e *Election) Init() {
	e.Lock()
	if e.inited {
		e.Unlock()
		e.logInfo("Only one init available")
		return
	}
	e.inited = true
	e.Unlock()
	for {
		if !e.isInit() {
			break
		}
		e.process()
		if !e.isInit() {
			break
		}
		wait(e.CheckTimeout)
	}
	e.logDebug("I'm finished")
}

func (e *Election) renew() error {
	_, _, err := e.Client.Session().Renew(e.sessionID, nil)
	return err
}

func (e *Election) destroySession(sesID string) error {
	_, err := e.Client.Session().Destroy(sesID, nil)
	if err != nil {
		e.logError("Destroy session error " + err.Error())
	}
	return err
}

func (e *Election) destroyCurrentSession() (err error) {
	if e.sessionID != "" {
		err = e.destroySession(e.sessionID)
		e.sessionID = ""
	}
	return
}

func (e *Election) needAcquire() bool {
	var (
		res string
		err error
	)
	for {
		res, err = e.getKvSession()
		if err != nil {
			e.disableLeader()
			wait(e.CheckTimeout)
		} else {
			break
		}
	}
	if e.sessionID != "" && e.sessionID == res {
		e.enableLeader(e.Master)
	}
	e.disableLeader()

	return res == ""
}

func (e *Election) process() {
	e.waitSession()
	if !e.leader {
		if !e.needAcquire() {
			return
		}
		e.logInfo("Try to acquire")
		res, err := e.acquire()
		if res && err == nil {
			e.enableLeader(e.Master)
		}
	} else {
		e.renew()
	}
}

func (e *Election) enableLeader(master string) {
	e.Lock()
	defer e.Unlock()
	if e.isInit() {
		e.leader = true
		e.logDebug("I'm a leader!")
		if e.Event != nil {
			e.Event.EventLeader(true, master)
			return
		}
	}
	e.Event.EventLeader(false, master)
}

// Stop election process
func (e *Election) Stop() {
	e.RLock()
	if !e.inited {
		e.RUnlock()
		return
	}
	e.RUnlock()
	e.stop <- struct{}{}
	<-e.success
}

func (e *Election) isInit() bool {
	for {
		select {
		case <-e.stop:
			e.inited = false
			e.logDebug("Stop signal received")
			e.disableLeader()
			e.destroyCurrentSession()
			e.success <- struct{}{}
			e.logDebug("Send success")
		default:
			return e.inited
		}
	}
}

func (e *Election) waitSession() {
	for {
		isset, err := e.checkSession()

		if isset {
			break
		}
		e.disableLeader()
		if err != nil {
			e.logDebug("Try to get session info again.")
			if !e.isInit() {
				break
			}
			wait(e.CheckTimeout)
			continue
		}
		err = e.createSession()

		if err == nil {
			e.logDebug("Session " + e.sessionID + " created")
			break
		}
		wait(e.CheckTimeout)
	}
}

func wait(t time.Duration) {
	runtime.Gosched()
	time.Sleep(t)
}

func (e *Election) logError(err string) {
	log.Error(e.LogPrefix + " [ERROR] " + err)
}

func (e *Election) logDebug(s string) {
	log.Debug(e.LogPrefix + " [DEBUG] " + s)
}

func (e *Election) logInfo(s string) {
	log.Info(e.LogPrefix + " [INFO] " + s)
}
