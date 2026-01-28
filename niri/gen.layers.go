package niri

import "github.com/frizinak/goniri/niri/types"

type layersRequest struct {
	r *layersResponse
}

func NewLayersRequest() layersRequest {
	return layersRequest{r: new(layersResponse)}
}

func (r layersRequest) MarshalJSON() ([]byte, error) {
	return []byte(`"Layers"`), nil
}

func (r layersRequest) response() Response {
	return r.r
}

func (r layersRequest) Response() []types.LayerSurface {
	return r.r.Layers
}
