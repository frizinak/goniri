package niri

import "github.com/frizinak/goniri/niri/types"

type workspacesRequest struct {
	r *workspacesResponse
}

func NewWorkspacesRequest() workspacesRequest {
	return workspacesRequest{r: new(workspacesResponse)}
}

func (r workspacesRequest) MarshalJSON() ([]byte, error) {
	return []byte(`"Workspaces"`), nil
}

func (r workspacesRequest) response() Response {
	return r.r
}

func (r workspacesRequest) Response() []types.Workspace {
	return r.r.Workspaces
}
