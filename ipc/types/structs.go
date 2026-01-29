package types

type Cast struct {
	StreamID      uint64     `json:"stream_id"`
	SessionID     uint64     `json:"session_id"`
	Kind          CastKind   `json:"kind"`
	Target        CastTarget `json:"target"`
	DynamicTarget bool       `json:"is_dynamic_target"`
	Active        bool       `json:"is_active"`
	PID           int        `json:"pid,omitempty"`
	PWNodeID      uint32     `json:"pw_node_id,omitempty"`
}

type ConfiguredMode struct {
	Width   uint16   `json:"width"`
	Height  uint16   `json:"height"`
	Refresh *float64 `json:"refresh"`
}

type ConfiguredPosition I32Position

type KeyboardLayouts struct {
	Names        []string `json:"names"`
	CurrentIndex uint8    `json:"current_idx"`
}

type LayerSurface struct {
	Namespace             string                            `json:"namespace"`
	Output                string                            `json:"output"`
	Layer                 Layer                             `json:"layer"`
	KeyboardInteractivity LayerSurfaceKeyboardInteractivity `json:"keyboard_interactivity"`
}

type LogicalOutput struct {
	I32Position
	Width     uint32    `json:"width"`
	Height    uint32    `json:"height"`
	Scale     float64   `json:"scale"`
	Transform Transform `json:"transform"`
}

type Mode struct {
	Width     uint16 `json:"width"`
	Height    uint16 `json:"height"`
	Refresh   uint32 `json:"refresh_rate"`
	Preferred bool   `json:"is_preferred"`
}

func (m Mode) ConfiguredMode() ConfiguredMode {
	var r *float64
	if m.Refresh != 0 {
		n := float64(m.Refresh) / 1000
		r = &n
	}

	return ConfiguredMode{
		Width:   m.Width,
		Height:  m.Height,
		Refresh: r,
	}
}

func (m Mode) ModeToSet() ModeToSet {
	return NewModeToSet(m.ConfiguredMode())
}

type Output struct {
	Name         string         `json:"name"`
	Make         string         `json:"make"`
	Model        string         `json:"model"`
	Serial       *string        `json:"serial,omitempty"`
	PhysicalSize *I32Size       `json:"physical_size,omitempty"`
	Modes        []Mode         `json:"modes"`
	CurrentMode  *uint          `json:"current_mode,omitempty"`
	CustomMode   bool           `json:"is_custom_mode"`
	VRRSupported bool           `json:"vrr_supported"`
	VRREnabled   bool           `json:"vrr_enabled"`
	Logical      *LogicalOutput `json:"logical,omitempty"`
}

type Overview struct {
	Open bool `json:"is_open"`
}

type PickedColor struct {
	Color RGB `json:"rgb"`
}

type Timestamp struct {
	Secs  uint64 `json:"secs"`
	Nanos uint32 `json:"nanos"`
}

type VRRToSet struct {
	VRR      bool `json:"vrr"`
	OnDemand bool `json:"on_demand"`
}

type Window struct {
	ID             uint64       `json:"id"`
	Title          *string      `json:"title,omitempty"`
	AppID          *string      `json:"app_id,omitempty"`
	PID            *int32       `json:"pid,omitempty"`
	WorkspaceID    *uint64      `json:"workspace_id,omitempty"`
	Focused        bool         `json:"is_focused"`
	Floating       bool         `json:"is_floating"`
	Urgent         bool         `json:"is_urgent"`
	Layout         WindowLayout `json:"layout"`
	FocusTimestamp *Timestamp   `json:"focus_timestamp,omitempty"`
}

type WindowLayout struct {
	PosInScrollingLayout   *Position    `json:"pos_in_scrolling_layout,omitempty"`
	TileSize               F64Size      `json:"tile_size"`
	WindowSize             I32Size      `json:"window_size"`
	TilePosInWorkspaceView *F64Position `json:"tile_pos_in_workspace_view,omitempty"`
	WindowOffsetInTile     F64Position  `json:"window_offset_in_tile"`
}

type Workspace struct {
	ID             uint64  `json:"id"`
	Index          uint8   `json:"idx"`
	Name           *string `json:"name,omitempty"`
	Output         *string `json:"output,omitempty"`
	Urgent         bool    `json:"is_urgent"`
	Active         bool    `json:"is_active"`
	Focused        bool    `json:"is_focused"`
	ActiveWindowID *uint64 `json:"active_window_id,omitempty"`
}
