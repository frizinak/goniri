package niri

type versionRequest struct {
	r *versionResponse
}

func NewVersionRequest() versionRequest {
	return versionRequest{r: new(versionResponse)}
}

func (r versionRequest) MarshalJSON() ([]byte, error) {
	return []byte(`"Version"`), nil
}

func (r versionRequest) response() Response {
	return r.r
}

func (r versionRequest) Response() string {
	return r.r.Version
}
