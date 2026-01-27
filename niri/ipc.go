package niri

import (
	"encoding/json"
	"errors"
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

type IPC struct {
	socket string

	conn io.Closer
	w    *json.Encoder
	r    *json.Decoder
}

func New(socket string) *IPC {
	return &IPC{socket: socket}
}

func (ipc *IPC) Init() error {
	conn, err := net.Dial("unix", ipc.socket)
	if err != nil {
		return err
	}
	ipc.conn = conn

	ipc.w = json.NewEncoder(conn)
	ipc.r = json.NewDecoder(conn)

	return nil
}

func (ipc *IPC) Do(req Request) error {
	if err := ipc.w.Encode(req); err != nil {
		return err
	}

	var r reply
	r.Ok = req.response()
	if err := ipc.r.Decode(&r); err != nil {
		return err
	}

	if r.Err != "" {
		return errors.New(r.Err)
	}

	return nil
}

func (ipc *IPC) Events(cb func(types.Event) error) error {
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

func (ipc *IPC) Close() error { return ipc.conn.Close() }
