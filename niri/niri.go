package niri

import (
	"context"
	"io"
	"net"
	"os"

	"github.com/frizinak/goniri/ipc"
	"github.com/frizinak/goniri/ipc/types"
)

type baseClient struct {
	c      ipc.Client
	closer io.Closer
}

func (c *baseClient) Close() error {
	if c.closer != nil {
		return c.closer.Close()
	}

	return nil
}

type Client struct {
	*baseClient
}

type EventsClient struct {
	*baseClient
}

func pv[T any](p *T) T {
	if p != nil {
		return *p
	}
	var z T
	return z
}

func (c *EventsClient) Events(ctx context.Context, h EventHandler) error {
	cb := func(e types.Event) error {
		switch {
		case nil != e.WorkspacesChanged:
			o := e.WorkspacesChanged.Workspaces
			p := make([]Workspace, len(o))
			for i, w := range o {
				p[i] = p[i].from(w)
			}
			h.Workspaces(p)

		case nil != e.WorkspaceUrgencyChanged:
			h.WorkspaceUrgency(
				WorkspaceID(e.WorkspaceUrgencyChanged.ID),
				e.WorkspaceUrgencyChanged.Urgent,
			)

		case nil != e.WorkspaceActivated:
			h.WorkspaceFocus(
				WorkspaceID(e.WorkspaceActivated.ID),
				e.WorkspaceActivated.Focused,
			)

		case nil != e.WorkspaceActiveWindowChanged:
			h.WorkspaceActiveWindow(
				WorkspaceID(e.WorkspaceActiveWindowChanged.WorkspaceID),
				WindowID(e.WorkspaceActiveWindowChanged.ActiveWindowID),
			)

		case nil != e.WindowsChanged:
			o := e.WindowsChanged.Windows
			p := make([]Window, len(o))
			for i, w := range o {
				p[i] = p[i].from(w)
			}
			h.Windows(p)

		case nil != e.WindowOpenedOrChanged:
			h.Window(
				Window{}.from(e.WindowOpenedOrChanged.Window),
			)

		case nil != e.WindowClosed:
			h.WindowClosed(WindowID(e.WindowClosed.ID))

		case nil != e.WindowFocusChanged:
			h.WindowFocus(WindowID(e.WindowFocusChanged.ID))

		case nil != e.WindowFocusTimestampChanged:
			h.WindowFocusTimestamp(
				WindowID(e.WindowFocusTimestampChanged.ID),
				stamp(e.WindowFocusTimestampChanged.FocusTimestamp),
			)

		case nil != e.WindowUrgencyChanged:
			h.WindowUrgency(
				WindowID(e.WindowUrgencyChanged.ID),
				e.WindowUrgencyChanged.Urgent,
			)

		case nil != e.WindowLayoutsChanged:
			o := e.WindowLayoutsChanged.Changes
			p := make(map[WindowID]WindowLayout, len(o))
			for _, w := range o {
				id := WindowID(w.ID)
				p[id] = p[id].from(w.WindowLayout)
			}
			h.WindowLayouts(p)

		case nil != e.KeyboardLayoutsChanged:
			h.KeyboardLayouts(
				int(e.KeyboardLayoutsChanged.KeyboardLayouts.CurrentIndex),
				e.KeyboardLayoutsChanged.KeyboardLayouts.Names,
			)

		case nil != e.KeyboardLayoutSwitched:
			h.KeyboardLayoutSwitched(int(e.KeyboardLayoutSwitched.Index))

		case nil != e.OverviewOpenedOrClosed:
			h.Overview(e.OverviewOpenedOrClosed.Open)

		case nil != e.ConfigLoaded:
			h.ConfigLoaded(!e.ConfigLoaded.Failed)

		case nil != e.ScreenshotCaptured:
			h.Screenshot(e.ScreenshotCaptured.Path)

		// TODO
		case nil != e.CastsChanged:
			var _ *types.CastsChanged
		case nil != e.CastStartedOrChanged:
			var _ *types.CastStartedOrChanged
		case nil != e.CastStopped:
			var _ *types.CastStopped
		}

		return nil
	}

	done := make(chan error, 1)
	go func() {
		done <- c.c.Events(cb)
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-done:
		return err
	}
}

func NewClient(c ipc.Client) *Client {
	return &Client{baseClient: &baseClient{c: c}}
}

func NewEventsClient(c ipc.Client) *EventsClient {
	return &EventsClient{baseClient: &baseClient{c: c}}
}

func DialSocket(socket string) (*Client, error) {
	conn, err := net.Dial("unix", socket)
	if err != nil {
		return nil, err
	}

	return &Client{baseClient: &baseClient{
		c:      ipc.New(conn),
		closer: conn,
	}}, nil
}

func DialEventsSocket(socket string) (*EventsClient, error) {
	conn, err := net.Dial("unix", socket)
	if err != nil {
		return nil, err
	}

	return &EventsClient{baseClient: &baseClient{
		c:      ipc.New(conn),
		closer: conn,
	}}, nil
}

func DialSocketDefault() (*Client, error) {
	return DialSocket(os.Getenv("NIRI_SOCKET"))
}

func DialEventsSocketDefault() (*EventsClient, error) {
	return DialEventsSocket(os.Getenv("NIRI_SOCKET"))
}
