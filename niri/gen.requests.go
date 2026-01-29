package niri

import "github.com/frizinak/goniri/niri/types"

type focusedWindowRequest struct {
	r *focusedWindowResponse
}

func RequestFocusedWindow() focusedWindowRequest {
	return focusedWindowRequest{r: new(focusedWindowResponse)}
}

func (r focusedWindowRequest) MarshalJSON() ([]byte, error) {
	return []byte(`"FocusedWindow"`), nil
}

func (r focusedWindowRequest) response() Response {
	return r.r
}

func (r focusedWindowRequest) Response() *types.Window {
	return r.r.FocusedWindow
}

type versionRequest struct {
	r *versionResponse
}

func RequestVersion() versionRequest {
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

type outputsRequest struct {
	r *outputsResponse
}

func RequestOutputs() outputsRequest {
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

type workspacesRequest struct {
	r *workspacesResponse
}

func RequestWorkspaces() workspacesRequest {
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

type windowsRequest struct {
	r *windowsResponse
}

func RequestWindows() windowsRequest {
	return windowsRequest{r: new(windowsResponse)}
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

type layersRequest struct {
	r *layersResponse
}

func RequestLayers() layersRequest {
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

type keyboardLayoutsRequest struct {
	r *keyboardLayoutsResponse
}

func RequestKeyboardLayouts() keyboardLayoutsRequest {
	return keyboardLayoutsRequest{r: new(keyboardLayoutsResponse)}
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

type focusedOutputRequest struct {
	r *focusedOutputResponse
}

func RequestFocusedOutput() focusedOutputRequest {
	return focusedOutputRequest{r: new(focusedOutputResponse)}
}

func (r focusedOutputRequest) MarshalJSON() ([]byte, error) {
	return []byte(`"FocusedOutput"`), nil
}

func (r focusedOutputRequest) response() Response {
	return r.r
}

func (r focusedOutputRequest) Response() *types.Output {
	return r.r.FocusedOutput
}

type pickWindowRequest struct {
	r *pickWindowResponse
}

func RequestPickWindow() pickWindowRequest {
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

type pickColorRequest struct {
	r *pickColorResponse
}

func RequestPickColor() pickColorRequest {
	return pickColorRequest{r: new(pickColorResponse)}
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

type overviewStateRequest struct {
	r *overviewStateResponse
}

func RequestOverviewState() overviewStateRequest {
	return overviewStateRequest{r: new(overviewStateResponse)}
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

type castsRequest struct {
	r *castsResponse
}

func RequestCasts() castsRequest {
	return castsRequest{r: new(castsResponse)}
}

func (r castsRequest) MarshalJSON() ([]byte, error) {
	return []byte(`"Casts"`), nil
}

func (r castsRequest) response() Response {
	return r.r
}

func (r castsRequest) Response() []types.Cast {
	return r.r.Casts
}
