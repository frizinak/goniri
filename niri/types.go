package niri

import (
	"encoding/json"
	"errors"
)

type Workspace struct {
}

type Window struct {
	ID             uint64       `json:"id"`
	Title          string       `json:"title"`
	AppID          string       `json:"app_id"`
	PID            int32        `json:"pid"`
	WorkspaceID    uint64       `json:"uint64"`
	Focused        bool         `json:"is_focused"`
	Floating       bool         `json:"is_floating"`
	Urgent         bool         `json:"is_urgent"`
	Layout         WindowLayout `json:"layout"`
	FocusTimestamp *Timestamp   `json:"focus_timestamp"`
}

type WindowLayout struct {
	PosInScrollingLayout   [2]uint64  `json:"pos_in_scrolling_layout"`
	TileSize               [2]float64 `json:"tile_size"`
	WindowSize             [2]int32   `json:"window_size"`
	TilePosInWorkspaceView [2]float64 `json:"tile_pos_in_workspace_view"`
	WindowOffsetInTile     [2]float64 `json:"window_offset_in_tile"`
}

type Timestamp struct {
	Secs  uint64 `json:"secs"`
	Nanos uint32 `json:"nanos"`
}

type KeyboardLayouts struct {
	Names        []string `json:"names"`
	CurrentIndex uint8    `json:"current_idx"`
}

type Cast struct {
	StreamID        uint64     `json:"stream_id"`
	SessionID       uint64     `json:"session_id"`
	Kind            CastKind   `json:"kind"`
	Target          CastTarget `json:"target"`
	IsDynamicTarget bool       `json:"is_dynamic_target"`
	IsActive        bool       `json:"is_active"`
	PID             int32      `json:"pid"`
	PWNodeID        uint32     `json:"pw_node_id"`
}

type CastKind string

type CastTarget struct {
	Output struct {
		Name string `json:"name"`
	} `json:"Output"`
	Window struct {
		ID uint64 `json:"id"`
	} `json:"Window"`
}

// ----------------------------------------------------------------------
// Event
// ----------------------------------------------------------------------

type Event struct {
	WorkspacesChanged            *WorkspacesChanged            `json:"WorkspacesChanged"`
	WorkspaceUrgencyChanged      *WorkspaceUrgencyChanged      `json:"WorkspaceUrgencyChanged"`
	WorkspaceActivated           *WorkspaceActivated           `json:"WorkspaceActivated"`
	WorkspaceActiveWindowChanged *WorkspaceActiveWindowChanged `json:"WorkspaceActiveWindowChanged"`
	WindowsChanged               *WindowsChanged               `json:"WindowsChanged"`
	WindowOpenedOrChanged        *WindowOpenedOrChanged        `json:"WindowOpenedOrChanged"`
	WindowClosed                 *WindowClosed                 `json:"WindowClosed"`
	WindowFocusChanged           *WindowFocusChanged           `json:"WindowFocusChanged"`
	WindowFocusTimestampChanged  *WindowFocusTimestampChanged  `json:"WindowFocusTimestampChanged"`
	WindowUrgencyChanged         *WindowUrgencyChanged         `json:"WindowUrgencyChanged"`
	WindowLayoutsChanged         *WindowLayoutsChanged         `json:"WindowLayoutsChanged"`
	KeyboardLayoutsChanged       *KeyboardLayoutsChanged       `json:"KeyboardLayoutsChanged"`
	KeyboardLayoutSwitched       *KeyboardLayoutSwitched       `json:"KeyboardLayoutSwitched"`
	OverviewOpenedOrClosed       *OverviewOpenedOrClosed       `json:"OverviewOpenedOrClosed"`
	ConfigLoaded                 *ConfigLoaded                 `json:"ConfigLoaded"`
	ScreenshotCaptured           *ScreenshotCaptured           `json:"ScreenshotCaptured"`
	CastsChanged                 *CastsChanged                 `json:"CastsChanged"`
	CastStartedOrChanged         *CastStartedOrChanged         `json:"CastStartedOrChanged"`
	CastStopped                  *CastStopped                  `json:"CastStopped"`
}

type WorkspacesChanged struct {
	Workspaces []Workspace `json:"workspace"`
}

type WorkspaceUrgencyChanged struct {
	ID     uint64 `json:"id"`
	Urgent bool   `json:"urgent"`
}

type WorkspaceActivated struct {
	ID      uint64 `json:"id"`
	Focused bool   `json:"focused"`
}

type WorkspaceActiveWindowChanged struct {
	WorkspaceID    uint64 `json:"workspace_id"`
	ActiveWindowID uint64 `json:"active_window_id"`
}

type WindowsChanged struct {
	Windows []Window `json:"windows"`
}

type WindowOpenedOrChanged struct {
	Window Window `json:"window"`
}

type WindowClosed struct {
	ID uint64 `json:"id"`
}

type WindowFocusChanged struct {
	ID uint64 `json:"id"`
}

type WindowFocusTimestampChanged struct {
	ID             uint64     `json:"id"`
	FocusTimestamp *Timestamp `json:"focus_timestamp"`
}

type WindowUrgencyChanged struct {
	ID     uint64 `json:"id"`
	Urgent bool   `json:"urgent"`
}

type WindowLayoutChange struct {
	ID uint64
	WindowLayout
}

type WindowLayoutsChanged struct {
	Changes []WindowLayoutChange `json:"-"`
}

func (w *WindowLayoutsChanged) UnmarshalJSON(d []byte) error {
	type raw struct {
		Plain [][]json.RawMessage `json:"changes"`
	}
	r := &raw{}
	if err := json.Unmarshal(d, r); err != nil {
		return err
	}

	w.Changes = make([]WindowLayoutChange, 0, len(r.Plain))
	for _, p := range r.Plain {
		if len(p) == 0 {
			continue
		}
		if len(p) != 2 {
			return errors.New("unexpected data structurce in WindowLayoutsChanged")
		}

		var c WindowLayoutChange
		if err := json.Unmarshal(p[0], &c.ID); err != nil {
			return err
		}
		if err := json.Unmarshal(p[1], &c.WindowLayout); err != nil {
			return err
		}

		w.Changes = append(w.Changes, c)
	}

	return nil
}

type KeyboardLayoutsChanged struct {
	KeyboardLayouts KeyboardLayouts `json:"keyboard_layouts"`
}

type KeyboardLayoutSwitched struct {
	Index uint8 `json:"idx"`
}

type OverviewOpenedOrClosed struct {
	Open bool `json:"is_open"`
}

type ConfigLoaded struct {
	Failed bool `json:"failed"`
}

type ScreenshotCaptured struct {
	Path string `json:"path"`
}

type CastsChanged struct {
	Casts []Cast `json:"casts"`
}

type CastStartedOrChanged struct {
	Cast Cast `json:"cast"`
}

type CastStopped struct {
	StreamID uint64 `json:"stream_id"`
}
