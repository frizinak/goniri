package niri

import "github.com/frizinak/goniri/niri/types"

type focusedOutputRequest struct {
	r *focusedOutputResponse
}

func NewFocusedOutputRequest() focusedOutputRequest {
	return focusedOutputRequest{r: new(focusedOutputResponse)}
}

func (r focusedOutputRequest) MarshalJSON() ([]byte, error) {
	return []byte(`"FocusedOutput"`), nil
}

func (r focusedOutputRequest) response() Response {
	return r.r
}

func (r focusedOutputRequest) Response() *types.Output {
	return r.r.FocusedOutput
}
