package types

import "encoding/json"

type CastKind string

const (
	CastKindPipeWire      CastKind = "PipeWire"
	CastKindWlrScreencopy CastKind = "WlrScreencopy"
)

type CastTarget struct {
	Kind       CastTargetKind
	OutputName string
	WindowID   uint64
}

func (t *CastTarget) UnmarshalJSON(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	for k, v := range raw {
		t.Kind = CastTargetKind(k)
		switch t.Kind {
		case CastTargetNothing:
			return nil
		case CastTargetOutput:
			var r struct {
				Name string `json:"name"`
			}
			err := json.Unmarshal(v, &r)
			t.OutputName = r.Name
			return err
		case CastTargetWindow:
			var r struct {
				ID uint64 `json:"id"`
			}
			err := json.Unmarshal(v, &r)
			t.WindowID = r.ID
			return err
		}
	}

	return nil
}

type ColumnDisplay string

const (
	ColumnDisplayNormal ColumnDisplay = "Normal"
	ColumnDisplayTabbed ColumnDisplay = "Tabbed"
)

type Event struct {
	WorkspacesChanged            *WorkspacesChanged            `json:"WorkspacesChanged,omitempty"`
	WorkspaceUrgencyChanged      *WorkspaceUrgencyChanged      `json:"WorkspaceUrgencyChanged,omitempty"`
	WorkspaceActivated           *WorkspaceActivated           `json:"WorkspaceActivated,omitempty"`
	WorkspaceActiveWindowChanged *WorkspaceActiveWindowChanged `json:"WorkspaceActiveWindowChanged,omitempty"`
	WindowsChanged               *WindowsChanged               `json:"WindowsChanged,omitempty"`
	WindowOpenedOrChanged        *WindowOpenedOrChanged        `json:"WindowOpenedOrChanged,omitempty"`
	WindowClosed                 *WindowClosed                 `json:"WindowClosed,omitempty"`
	WindowFocusChanged           *WindowFocusChanged           `json:"WindowFocusChanged,omitempty"`
	WindowFocusTimestampChanged  *WindowFocusTimestampChanged  `json:"WindowFocusTimestampChanged,omitempty"`
	WindowUrgencyChanged         *WindowUrgencyChanged         `json:"WindowUrgencyChanged,omitempty"`
	WindowLayoutsChanged         *WindowLayoutsChanged         `json:"WindowLayoutsChanged,omitempty"`
	KeyboardLayoutsChanged       *KeyboardLayoutsChanged       `json:"KeyboardLayoutsChanged,omitempty"`
	KeyboardLayoutSwitched       *KeyboardLayoutSwitched       `json:"KeyboardLayoutSwitched,omitempty"`
	OverviewOpenedOrClosed       *OverviewOpenedOrClosed       `json:"OverviewOpenedOrClosed,omitempty"`
	ConfigLoaded                 *ConfigLoaded                 `json:"ConfigLoaded,omitempty"`
	ScreenshotCaptured           *ScreenshotCaptured           `json:"ScreenshotCaptured,omitempty"`
	CastsChanged                 *CastsChanged                 `json:"CastsChanged,omitempty"`
	CastStartedOrChanged         *CastStartedOrChanged         `json:"CastStartedOrChanged,omitempty"`
	CastStopped                  *CastStopped                  `json:"CastStopped,omitempty"`
}

type HSyncPolarity string

const (
	HSyncPolarityPHSync HSyncPolarity = "PHSync"
	HSyncPolarityNHSync HSyncPolarity = "NHSync"
)

type Layer string

const (
	LayerBackground Layer = "Background"
	LayerBottom     Layer = "Bottom"
	LayerTop        Layer = "Top"
	LayerOverlay    Layer = "Overlay"
)

type LayerSurfaceKeyboardInteractivity string

const (
	LayerSurfaceKeyboardInteractivityNone      LayerSurfaceKeyboardInteractivity = "None"
	LayerSurfaceKeyboardInteractivityExclusive LayerSurfaceKeyboardInteractivity = "Exclusive"
	LayerSurfaceKeyboardInteractivityOnDemand  LayerSurfaceKeyboardInteractivity = "OnDemand"
)

type LayoutSwitchTargetKind string

const (
	LayoutSwitchTargetNext  LayoutSwitchTargetKind = "Next"
	LayoutSwitchTargetPrev  LayoutSwitchTargetKind = "Prev"
	LayoutSwitchTargetIndex LayoutSwitchTargetKind = "Index"
)

type LayoutSwitchTarget struct {
	Kind  LayoutSwitchTargetKind
	Index *uint8
}

func (t *LayoutSwitchTarget) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		t.Kind = LayoutSwitchTargetKind(s)
		return nil
	}

	var o map[string]uint8
	if err := json.Unmarshal(data, &o); err != nil {
		return err
	}

	for k, v := range o {
		if k == "Index" {
			t.Kind = LayoutSwitchTargetIndex
			t.Index = &v
		}
	}

	return nil
}

type ModeToSet struct {
	Kind ModeToSetKind
	ConfiguredMode
}

func (m *ModeToSet) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		m.Kind = ModeToSetKind(s)
		return nil
	}

	var o map[string]ConfiguredMode
	if err := json.Unmarshal(data, &o); err != nil {
		return err
	}

	for k, v := range o {
		if k == string(ModeToSetSpecific) {
			m.Kind = ModeToSetSpecific
			m.ConfiguredMode = v
			break
		}
	}

	return nil
}

type OutputConfigChanged string

const (
	OutputConfigChangedApplied          OutputConfigChanged = "Applied"
	OutputConfigChangedOutputWasMissing OutputConfigChanged = "OutputWasMissing"
)

type PositionChange struct {
	Kind  PositionChangeKind
	Value float64
}

func (p *PositionChange) UnmarshalJSON(data []byte) error {
	var o map[string]float64
	if err := json.Unmarshal(data, &o); err != nil {
		return err
	}

	for k, v := range o {
		p.Kind = PositionChangeKind(k)
		p.Value = v
		break
	}

	return nil
}

type PositionToSet struct {
	Kind PositionToSetKind
	ConfiguredPosition
}

func (p *PositionToSet) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		p.Kind = PositionToSetKind(s)
		return nil
	}

	var o map[string]ConfiguredPosition
	if err := json.Unmarshal(data, &o); err != nil {
		return err
	}

	for k, v := range o {
		if k == string(PositionToSetSpecific) {
			p.Kind = PositionToSetSpecific
			p.ConfiguredPosition = v
			break
		}
	}

	return nil
}

type ScaleToSet struct {
	Kind  ScaleToSetKind
	Scale float64
}

func (p *ScaleToSet) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		p.Kind = ScaleToSetKind(s)
		return nil
	}

	var o map[string]float64
	if err := json.Unmarshal(data, &o); err != nil {
		return err
	}

	for k, v := range o {
		if k == string(ScaleToSetSpecific) {
			p.Kind = ScaleToSetSpecific
			p.Scale = v
			break
		}
	}

	return nil
}

type SizeChange struct {
	Kind       SizeChangeKind
	Fixed      int32
	Proportion float64
}

func (s *SizeChange) UnmarshalJSON(data []byte) error {
	var o map[string]json.RawMessage
	if err := json.Unmarshal(data, &o); err != nil {
		return err
	}

	for k, v := range o {
		s.Kind = SizeChangeKind(k)
		switch s.Kind {
		case SizeChangeSetFixed, SizeChangeAdjustFixed:
			var r int32
			err := json.Unmarshal(v, &r)
			s.Fixed = r
			return err
		case SizeChangeSetProportion, SizeChangeAdjustProportion:
			var r float64
			err := json.Unmarshal(v, &r)
			s.Proportion = r
			return err
		}
		break
	}

	return nil
}

type Transform string

const (
	TransformNormal     Transform = "Normal"
	Transform90         Transform = "_90"
	Transform180        Transform = "_180"
	Transform270        Transform = "_270"
	TransformFlipped    Transform = "Flipped"
	TransformFlipped90  Transform = "Flipped90"
	TransformFlipped180 Transform = "Flipped180"
	TransformFlipped270 Transform = "Flipped270"
)

type VSyncPolarity string

const (
	VSyncPolarityPVSync VSyncPolarity = "PVSync"
	VSyncPolarityNVSync VSyncPolarity = "NVSync"
)

type WorkspaceReferenceArg struct {
	Kind  WorkspaceReferenceArgKind
	ID    uint64
	Index uint8
	Name  string
}

func (w *WorkspaceReferenceArg) UnmarshalJSON(data []byte) error {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	for k, v := range raw {
		w.Kind = WorkspaceReferenceArgKind(k)
		switch w.Kind {
		case WorkspaceReferenceArgID:
			var r uint64
			err := json.Unmarshal(v, &r)
			w.ID = r
			return err
		case WorkspaceReferenceArgIndex:
			var r uint8
			err := json.Unmarshal(v, &r)
			w.Index = r
			return err
		case WorkspaceReferenceArgName:
			var r string
			err := json.Unmarshal(v, &r)
			w.Name = r
			return err
		}
	}

	return nil
}
