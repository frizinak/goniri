package niri

import "github.com/frizinak/goniri/niri/types"

type outputsRequest struct {
	r *outputsResponse
}

func NewOutputsRequest() outputsRequest {
	return outputsRequest{r: new(outputsResponse)}
}

func (r outputsRequest) MarshalJSON() ([]byte, error) {
	return []byte(`"Outputs"`), nil
}

func (r outputsRequest) response() Response {
	return r.r
}

func (r outputsRequest) Response() map[string]types.Output {
	return r.r.Outputs
}
