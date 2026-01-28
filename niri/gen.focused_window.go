package niri

import "github.com/frizinak/goniri/niri/types"

type focusedWindowRequest struct {
	r *focusedWindowResponse
}

func NewFocusedWindowRequest() focusedWindowRequest {
	return focusedWindowRequest{r: new(focusedWindowResponse)}
}

func (r focusedWindowRequest) MarshalJSON() ([]byte, error) {
	return []byte(`"FocusedWindow"`), nil
}

func (r focusedWindowRequest) response() Response {
	return r.r
}

func (r focusedWindowRequest) Response() *types.Window {
	return r.r.FocusedWindow
}
