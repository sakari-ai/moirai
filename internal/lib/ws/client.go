package ws

import (
	"errors"
	"time"

	"github.com/gorilla/websocket"
)

type client struct {
	conn *websocket.Conn
}

type SocketClient interface {
	Send(msg interface{}) error
	SendBytes(msg []byte) error
	Close() error
}

func NewClient(address string) (*client, error) {
	c := new(client)
	err := c.connect(address)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *client) connect(address string) error {
	conn, _, err := websocket.DefaultDialer.Dial(address, nil)
	if err != nil {
		return err
	}
	c.conn = conn
	return nil
}

func (c *client) Send(msg interface{}) error {
	if c.conn == nil {
		return errors.New("connection is lost")
	}
	return c.conn.WriteJSON(msg)
}

func (c *client) SendBytes(msg []byte) error {
	if c.conn == nil {
		return errors.New("connection is lost")
	}
	return c.conn.WriteMessage(websocket.BinaryMessage, msg)
}

func (c *client) Close() error {
	err := c.conn.WriteControl(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""), time.Now().Add(time.Second))
	if err != nil {
		return err
	}
	err = c.conn.Close()
	return err
}
