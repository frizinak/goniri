package niri

import "github.com/frizinak/goniri/niri/types"

type windowsRequest struct {
	r *windowsResponse
}

func NewWindowsRequest() windowsRequest {
	return windowsRequest{r: &windowsResponse{}}
}

func (r windowsRequest) MarshalJSON() ([]byte, error) {
	return []byte(`"Windows"`), nil
}

func (r windowsRequest) response() Response {
	return r.r
}

func (r windowsRequest) Response() []types.Window {
	return r.r.Windows
}
