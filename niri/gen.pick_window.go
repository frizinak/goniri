package niri

import "github.com/frizinak/goniri/niri/types"

type pickWindowRequest struct {
	r *pickWindowResponse
}

func NewPickWindowRequest() pickWindowRequest {
	return pickWindowRequest{r: new(pickWindowResponse)}
}

func (r pickWindowRequest) MarshalJSON() ([]byte, error) {
	return []byte(`"PickWindow"`), nil
}

func (r pickWindowRequest) response() Response {
	return r.r
}

func (r pickWindowRequest) Response() *types.Window {
	return r.r.PickedWindow
}
