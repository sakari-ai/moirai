package ws

import (
	"github.com/sakari-ai/moirai/log"
	"github.com/sakari-ai/moirai/log/field"
)

type socketPublisher struct {
	SocketClient
	address string
}

func NewSocketClient(address string) SocketClient {
	return &socketPublisher{address: address}
}

func (c socketPublisher) SendBytes(msg []byte) error {
	return c.sendToWSServer(msg)
}

func (c socketPublisher) connect() SocketClient {
	cl, _ := NewClient(c.address)
	c.SocketClient = cl
	return c.SocketClient
}

func (c socketPublisher) sendToWSServer(data []byte) error {
	var err error
	cl := c.SocketClient
	if cl == nil {
		cl = c.connect()
	}
	attempt := 0
	for {
		if attempt == 5 {
			break
		}
		if cl != nil {
			err = cl.SendBytes(data)
			if err == nil {
				break
			}
		}
		log.Error("can not send message to WS", field.Error(err))
		if err != nil && cl != nil {
			_ = cl.Close()
		}
		cl = c.connect()
		attempt++
		continue
	}
	return err
}
