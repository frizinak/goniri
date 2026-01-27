package niri

import "github.com/frizinak/goniri/niri/types"

type eventStreamResponse string

type eventStreamRequest struct {
	r *eventStreamResponse
}

func newEventStreamRequest() eventStreamRequest {
	var n eventStreamResponse
	return eventStreamRequest{r: &n}
}

func (r eventStreamRequest) MarshalJSON() ([]byte, error) {
	return []byte(`"EventStream"`), nil
}

func (r eventStreamRequest) response() Response {
	return r.r

}

//go:generate go run generate/request.go focusedWindow FocusedWindow *types.Window
type focusedWindowResponse struct {
	FocusedWindow *types.Window `json:"FocusedWindow"`
}

//go:generate go run generate/request.go version Version string
type versionResponse struct {
	Version string `json:"Version"`
}

//go:generate go run generate/request.go outputs Outputs map[string]types.Output
type outputsResponse struct {
	Outputs map[string]types.Output `json:"Outputs"`
}

//go:generate go run generate/request.go workspaces Workspaces []types.Workspace
type workspacesResponse struct {
	Workspaces []types.Workspace `json:"Workspaces"`
}

//go:generate go run generate/request.go windows Windows []types.Window
type windowsResponse struct {
	Windows []types.Window `json:"Windows"`
}

//go:generate go run generate/request.go layers Layers []types.LayerSurface
type layersResponse struct {
	Layers []types.LayerSurface `json:"Layers"`
}

//go:generate go run generate/request.go keyboardLayouts KeyboardLayouts types.KeyboardLayouts
type keyboardLayoutsResponse struct {
	KeyboardLayouts types.KeyboardLayouts `json:"KeyboardLayouts"`
}

//go:generate go run generate/request.go focusedOutput FocusedOutput *types.Output
type focusedOutputResponse struct {
	FocusedOutput *types.Output `json:"FocusedOutput"`
}

//go:generate go run generate/request.go pickWindow PickedWindow *types.Window
type pickWindowResponse struct {
	PickedWindow *types.Window `json:"PickedWindow"`
}

//go:generate go run generate/request.go pickColor PickedColor *types.PickedColor
type pickColorResponse struct {
	PickedColor *types.PickedColor `json:"PickedColor"`
}

//go:generate go run generate/request.go overviewState OverviewState *types.Overview
type overviewStateResponse struct {
	OverviewState *types.Overview `json:"OverviewState"`
	types.OverviewOpenedOrClosed
}

// Documented, but doens't work.
// -go:generate go run generate/request.go casts Casts []types.Cast
// type castsResponse struct {
// 	Casts []types.Cast `json:"Casts"`
// }

// TODO OutputAction
// TODO Action
