package niri

// ----------------------------------------------------------------------
// EventStream
// ----------------------------------------------------------------------

type eventStreamRequest struct {
	r *eventStreamResponse
}

func newEventStreamRequest() eventStreamRequest {
	var n eventStreamResponse
	return eventStreamRequest{r: &n}
}

func (r eventStreamRequest) MarshalJSON() ([]byte, error) {
	return []byte(`"EventStream"`), nil
}

func (r eventStreamRequest) response() Response {
	return r.r
}

type eventStreamResponse string

// ----------------------------------------------------------------------
// FocusedWindow
// ----------------------------------------------------------------------

type focusedWindowRequest struct {
	r *FocusedWindowResponse
}

func NewFocusedWindowRequest() focusedWindowRequest {
	return focusedWindowRequest{r: &FocusedWindowResponse{}}
}

func (r focusedWindowRequest) MarshalJSON() ([]byte, error) {
	return []byte(`"FocusedWindow"`), nil
}

func (r focusedWindowRequest) response() Response {
	return r.r
}

func (r focusedWindowRequest) Response() Window {
	return r.r.FocusedWindow
}

type FocusedWindowResponse struct {
	FocusedWindow Window `json:"FocusedWindow"`
}
