package types

import (
	"encoding/json"
	"errors"
)

type CastTargetKind string

const (
	CastTargetNothing CastTargetKind = "Nothing"
	CastTargetOutput  CastTargetKind = "Output"
	CastTargetWindow  CastTargetKind = "Window"
)

type LayoutSwitchTargetKind string

const (
	LayoutSwitchTargetNext  LayoutSwitchTargetKind = "Next"
	LayoutSwitchTargetPrev  LayoutSwitchTargetKind = "Prev"
	LayoutSwitchTargetIndex LayoutSwitchTargetKind = "Index"
)

type ModeToSetKind string

const (
	ModeToSetAutomatic ModeToSetKind = "Automatic"
	ModeToSetSpecific  ModeToSetKind = "Specific"
)

type PositionChangeKind string

const (
	PositionChangeSetFixed         PositionChangeKind = "SetFixed"
	PositionChangeSetProportion    PositionChangeKind = "SetProportion"
	PositionChangeAdjustFixed      PositionChangeKind = "AdjustFixed"
	PositionChangeAdjustProportion PositionChangeKind = "AdjustProportion"
)

type PositionToSetKind string

const (
	PositionToSetAutomatic PositionToSetKind = "Automatic"
	PositionToSetSpecific  PositionToSetKind = "Specific"
)

type ScaleToSetKind string

const (
	ScaleToSetAutomatic ScaleToSetKind = "Automatic"
	ScaleToSetSpecific  ScaleToSetKind = "Specific"
)

type SizeChangeKind string

const (
	SizeChangeSetFixed         SizeChangeKind = "SetFixed"
	SizeChangeSetProportion    SizeChangeKind = "SetProportion"
	SizeChangeAdjustFixed      SizeChangeKind = "AdjustFixed"
	SizeChangeAdjustProportion SizeChangeKind = "AdjustProportion"
)

type WorkspaceRefKind string

const (
	WorkspaceRefID    WorkspaceRefKind = "Id"
	WorkspaceRefIndex WorkspaceRefKind = "Index"
	WorkspaceRefName  WorkspaceRefKind = "Name"
)

type Position struct {
	Column int
	Tile   int
}

func (p *Position) UnmarshalJSON(data []byte) error {
	var raw [2]int
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	p.Column = raw[0]
	p.Tile = raw[1]
	return nil
}

type I32Position struct {
	X int32 `json:"x"`
	Y int32 `json:"y"`
}

type F64Position struct {
	X float64
	Y float64
}

func (p *F64Position) UnmarshalJSON(data []byte) error {
	var raw [2]float64
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	p.X = raw[0]
	p.Y = raw[1]
	return nil
}

type F64Size struct {
	Width  float64
	Height float64
}

func (s *F64Size) UnmarshalJSON(data []byte) error {
	var raw [2]float64
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	s.Width = raw[0]
	s.Height = raw[1]
	return nil
}

type I32Size struct {
	Width  int32
	Height int32
}

func (s *I32Size) UnmarshalJSON(data []byte) error {
	var raw [2]int32
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	s.Width = raw[0]
	s.Height = raw[1]
	return nil
}

type RGB struct {
	R, G, B float64
}

func (c *RGB) UnmarshalJSON(data []byte) error {
	var raw [3]float64
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	c.R = raw[0]
	c.G = raw[1]
	c.B = raw[2]
	return nil
}

// ----------------------------------------------------------------------
// Event
// ----------------------------------------------------------------------

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
