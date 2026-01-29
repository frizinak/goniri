package ipc

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"

	"github.com/frizinak/goniri/ipc/types"
)

type reply struct {
	Ok  Response `json:"Ok"`
	Err string   `json:"Err"`
}

type Response interface{}
type ResponseWithError interface {
	Err() error
}

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
		return fmt.Errorf("niri json error: %w", err)
	}

	if r.Err != "" {
		return fmt.Errorf("niri error: %w", errors.New(r.Err))
	}

	if we, ok := r.Ok.(ResponseWithError); ok {
		if err := we.Err(); err != nil {
			return fmt.Errorf("niri error: %w", err)
		}
	}

	return nil
}

func (ipc *ipc) Events(cb EventHandler) error {
	if err := ipc.Do(newEventStream()); err != nil {
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
