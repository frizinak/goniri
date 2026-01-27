package niri

import "github.com/frizinak/goniri/niri/types"

type keyboardLayoutsRequest struct {
	r *keyboardLayoutsResponse
}

func NewKeyboardLayoutsRequest() keyboardLayoutsRequest {
	return keyboardLayoutsRequest{r: &keyboardLayoutsResponse{}}
}

func (r keyboardLayoutsRequest) MarshalJSON() ([]byte, error) {
	return []byte(`"KeyboardLayouts"`), nil
}

func (r keyboardLayoutsRequest) response() Response {
	return r.r
}

func (r keyboardLayoutsRequest) Response() types.KeyboardLayouts {
	return r.r.KeyboardLayouts
}
