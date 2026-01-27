package niri

import "github.com/frizinak/goniri/niri/types"

type overviewStateRequest struct {
	r *overviewStateResponse
}

func NewOverviewStateRequest() overviewStateRequest {
	return overviewStateRequest{r: &overviewStateResponse{}}
}

func (r overviewStateRequest) MarshalJSON() ([]byte, error) {
	return []byte(`"OverviewState"`), nil
}

func (r overviewStateRequest) response() Response {
	return r.r
}

func (r overviewStateRequest) Response() *types.Overview {
	return r.r.OverviewState
}
