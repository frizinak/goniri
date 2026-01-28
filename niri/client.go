package niri

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"

	"github.com/frizinak/goniri/niri/types"
)

type reply struct {
	Ok  Response `json:"Ok"`
	Err string   `json:"Err"`
}

type Response interface{}

type Request interface {
	json.Marshaler
	response() Response
}

type EventHandler func(types.Event) error

type Client interface {
	Connect() error
	Do(Request) error
	Events(EventHandler) error
	io.Closer
}

type ipc struct {
	socket string

	conn io.Closer
	w    *json.Encoder
	r    *json.Decoder
}

func New(socket string) Client {
	return &ipc{socket: socket}
}

func (ipc *ipc) Connect() error {
	conn, err := net.Dial("unix", ipc.socket)
	if err != nil {
		return err
	}
	ipc.conn = conn

	ipc.w = json.NewEncoder(conn)
	ipc.r = json.NewDecoder(conn)

	return nil
}

func (ipc *ipc) Do(req Request) error {
	if err := ipc.w.Encode(req); err != nil {
		return err
	}

	var r reply
	r.Ok = req.response()
	if err := ipc.r.Decode(&r); err != nil {
		return fmt.Errorf("niri response json error: %w", err)
	}

	if r.Err != "" {
		return fmt.Errorf("niri replied with error: %w", errors.New(r.Err))
	}

	return nil
}

func (ipc *ipc) Events(cb EventHandler) error {
	if err := ipc.Do(newEventStreamRequest()); err != nil {
		return err
	}

	for {
		var ev types.Event
		if err := ipc.r.Decode(&ev); err != nil {
			return err
		}
		if err := cb(ev); err != nil {
			return err
		}
	}
}

func (ipc *ipc) Close() error { return ipc.conn.Close() }
