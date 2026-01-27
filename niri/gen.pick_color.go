package niri

import "github.com/frizinak/goniri/niri/types"

type pickColorRequest struct {
	r *pickColorResponse
}

func NewPickColorRequest() pickColorRequest {
	return pickColorRequest{r: &pickColorResponse{}}
}

func (r pickColorRequest) MarshalJSON() ([]byte, error) {
	return []byte(`"PickColor"`), nil
}

func (r pickColorRequest) response() Response {
	return r.r
}

func (r pickColorRequest) Response() *types.PickedColor {
	return r.r.PickedColor
}
