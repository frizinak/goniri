package niri

import (
	"time"

	"github.com/frizinak/goniri/ipc/types"
)

type WorkspaceID uint64
type WindowID uint64
type Index uint8

type Workspace struct {
	ID           WorkspaceID
	Index        Index
	Name         string
	Output       string
	Urgent       bool
	Active       bool
	Focused      bool
	ActiveWindow WindowID
}

func (w Workspace) from(o types.Workspace) Workspace {
	w.ID = WorkspaceID(o.ID)
	w.Index = Index(o.Index)
	w.Name = pv(o.Name)
	w.Output = pv(o.Output)
	w.Urgent = o.Urgent
	w.Active = o.Active
	w.Focused = o.Focused
	w.ActiveWindow = WindowID(pv(o.ActiveWindowID))
	return w
}

type Window struct {
	ID             WindowID
	Title          string
	AppID          string
	PID            int32
	WorkspaceID    WorkspaceID
	Focused        bool
	Floating       bool
	Urgent         bool
	Layout         WindowLayout
	FocusTimestamp time.Duration
}

func (w Window) from(o types.Window) Window {
	w.ID = WindowID(o.ID)
	w.Title = pv(o.Title)
	w.AppID = pv(o.AppID)
	w.PID = pv(o.PID)
	w.WorkspaceID = WorkspaceID(pv(o.WorkspaceID))
	w.Focused = o.Focused
	w.Floating = o.Floating
	w.Urgent = o.Urgent
	w.Layout = w.Layout.from(o.Layout)
	w.FocusTimestamp = stamp(o.FocusTimestamp)

	return w
}

type WindowLayout struct {
	X, Y int // 1-based

	OuterWidth  float64
	OuterHeight float64

	Width  int32
	Height int32

	OffsetX float64
	OffsetY float64
}

func (l WindowLayout) from(o types.WindowLayout) WindowLayout {
	if o.PosInScrollingLayout != nil {
		l.X = o.PosInScrollingLayout.Column
		l.Y = o.PosInScrollingLayout.Tile
	}

	l.OuterWidth = o.TileSize.Width
	l.OuterHeight = o.TileSize.Height

	// TODO?
	//if o.TilePosInWorkspaceView != nil {
	//}

	l.Width = o.WindowSize.Width
	l.Height = o.WindowSize.Height

	l.OffsetX = o.WindowOffsetInTile.X
	l.OffsetY = o.WindowOffsetInTile.Y

	return l
}

func stamp(ts *types.Timestamp) (dur time.Duration) {
	if ts != nil {
		dur = time.Second*time.Duration(ts.Secs) +
			time.Nanosecond*time.Duration(ts.Nanos)
	}
	return
}
