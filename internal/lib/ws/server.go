package ws

import (
	"context"
	"net/http"
	_ "net/http/pprof"
	"reflect"
	"sync"
	"syscall"

	"github.com/sakari-ai/moirai/log"
	"github.com/sakari-ai/moirai/log/field"
	http2 "github.com/sakari-ai/moirai/server/http"
	"github.com/gorilla/websocket"
	"golang.org/x/sys/unix"
)

var (
	sv *server
)

type Controller interface {
	Handle(msg []byte)
}

type server struct {
	epoller *epoll
	ctl     Controller
}

type EPollerServer interface {
	Start(context.Context, string)
}

func NewServer(ctl Controller) *server {
	sv = &server{
		ctl: ctl,
	}
	return sv
}

func (s *server) Start(ctx context.Context, host string) error {
	var rLimit syscall.Rlimit
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		return err
	}
	rLimit.Cur = rLimit.Max
	if err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		return err
	}

	// Enable pprof hooks
	//go func() {
	//	if err := http.ListenAndServe(host, nil); err != nil {
	//		log.Error("pprof failed", field.Error(err))
	//	}
	//}()

	// Start epoll
	var err error
	s.epoller, err = MkEpoll()
	if err != nil {
		return err
	}

	go start()

	http.HandleFunc("/v1/version", http2.VersionHandler)
	http.HandleFunc("/", wsHandler)
	if err := http.ListenAndServe(host, nil); err != nil {
		log.Error("Failed to enable handler", field.Error(err))
		return err
	}
	return nil
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade connection
	upgrader := websocket.Upgrader{}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	if err := sv.epoller.Add(conn); err != nil {
		log.Error("Failed to add connection", field.Error(err))
		_ = conn.Close()
	}
}

func start() {
	for {
		connections, err := sv.epoller.Wait()
		if err != nil {
			log.Error("Failed to epoll wait ", field.Error(err))
			continue
		}
		for _, conn := range connections {
			if conn == nil {
				break
			}
			_, msg, err := conn.ReadMessage()
			if err != nil {
				if err := sv.epoller.Remove(conn); err != nil {
					log.Error("Failed to remove connection", field.Error(err))
				}
				_ = conn.Close()
			} else {
				sv.ctl.Handle(msg)
			}
		}
	}
}

type epoll struct {
	fd          int
	connections map[int]*websocket.Conn
	lock        *sync.RWMutex
}

func MkEpoll() (*epoll, error) {
	fd, err := unix.EpollCreate1(0)
	if err != nil {
		return nil, err
	}
	return &epoll{
		fd:          fd,
		lock:        &sync.RWMutex{},
		connections: make(map[int]*websocket.Conn),
	}, nil
}

func (e *epoll) Add(conn *websocket.Conn) error {
	fd := websocketFD(conn)
	err := unix.EpollCtl(e.fd, syscall.EPOLL_CTL_ADD, fd, &unix.EpollEvent{Events: unix.POLLIN | unix.POLLHUP, Fd: int32(fd)})
	if err != nil {
		return err
	}
	e.lock.Lock()
	defer e.lock.Unlock()
	e.connections[fd] = conn
	if len(e.connections)%100 == 0 {
		log.Info("Total number of connections: %v", field.Int("connections", len(e.connections)))
	}
	return nil
}

func (e *epoll) Remove(conn *websocket.Conn) error {
	fd := websocketFD(conn)
	err := unix.EpollCtl(e.fd, syscall.EPOLL_CTL_DEL, fd, nil)
	if err != nil {
		return err
	}
	e.lock.Lock()
	defer e.lock.Unlock()
	delete(e.connections, fd)
	if len(e.connections)%100 == 0 {
		log.Info("Total number of connections: %v", field.Int("connection", len(e.connections)))
	}
	return nil
}

func (e *epoll) Wait() ([]*websocket.Conn, error) {
	events := make([]unix.EpollEvent, 100)
	n, err := unix.EpollWait(e.fd, events, 100)
	if err != nil {
		return nil, err
	}
	e.lock.RLock()
	defer e.lock.RUnlock()
	var connections []*websocket.Conn
	for i := 0; i < n; i++ {
		conn := e.connections[int(events[i].Fd)]
		connections = append(connections, conn)
	}
	return connections, nil
}

func websocketFD(conn *websocket.Conn) int {
	connVal := reflect.Indirect(reflect.ValueOf(conn)).FieldByName("conn").Elem()
	tcpConn := reflect.Indirect(connVal).FieldByName("conn")
	fdVal := tcpConn.FieldByName("fd")
	pfdVal := reflect.Indirect(fdVal).FieldByName("pfd")
	return int(pfdVal.FieldByName("Sysfd").Int())
}
