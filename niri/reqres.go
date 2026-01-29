package niri

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/frizinak/goniri/niri/types"
)

//go:generate go build -o generate/bin/request ./generate/request
//go:generate ./generate/bin/request create

type eventStreamResponse string

type eventStreamRequest struct {
	r *eventStreamResponse
}

func newEventStream() eventStreamRequest {
	return eventStreamRequest{r: new(eventStreamResponse)}
}

func (r eventStreamRequest) MarshalJSON() ([]byte, error) {
	return []byte(`"EventStream"`), nil
}

func (r eventStreamRequest) response() Response { return r.r }

//go:generate ./generate/bin/request focusedWindow FocusedWindow *types.Window
type focusedWindowResponse struct {
	FocusedWindow *types.Window `json:"FocusedWindow"`
}

//go:generate ./generate/bin/request version Version string
type versionResponse struct {
	Version string `json:"Version"`
}

//go:generate ./generate/bin/request outputs Outputs map[string]types.Output
type outputsResponse struct {
	Outputs map[string]types.Output `json:"Outputs"`
}

//go:generate ./generate/bin/request workspaces Workspaces []types.Workspace
type workspacesResponse struct {
	Workspaces []types.Workspace `json:"Workspaces"`
}

//go:generate ./generate/bin/request windows Windows []types.Window
type windowsResponse struct {
	Windows []types.Window `json:"Windows"`
}

//go:generate ./generate/bin/request layers Layers []types.LayerSurface
type layersResponse struct {
	Layers []types.LayerSurface `json:"Layers"`
}

//go:generate ./generate/bin/request keyboardLayouts KeyboardLayouts types.KeyboardLayouts
type keyboardLayoutsResponse struct {
	KeyboardLayouts types.KeyboardLayouts `json:"KeyboardLayouts"`
}

//go:generate ./generate/bin/request focusedOutput FocusedOutput *types.Output
type focusedOutputResponse struct {
	FocusedOutput *types.Output `json:"FocusedOutput"`
}

//go:generate ./generate/bin/request pickWindow PickedWindow *types.Window
type pickWindowResponse struct {
	PickedWindow *types.Window `json:"PickedWindow"`
}

//go:generate ./generate/bin/request pickColor PickedColor *types.PickedColor
type pickColorResponse struct {
	PickedColor *types.PickedColor `json:"PickedColor"`
}

type OutputMissingError struct{}

func (o OutputMissingError) Error() string {
	return string(types.OutputWasMissing)
}

type outputConfigChangedResponse struct {
	OutputConfigChanged types.OutputConfigChanged `json:"OutputConfigChanged"`
}

func (o outputConfigChangedResponse) Err() (err error) {
	if o.OutputConfigChanged == types.OutputWasMissing {
		err = OutputMissingError{}
		return
	}
	if o.OutputConfigChanged != types.OutputChangeApplied {
		err = errors.New(string(o.OutputConfigChanged))
	}

	return
}

//go:generate ./generate/bin/request overviewState OverviewState *types.Overview
type overviewStateResponse struct {
	OverviewState *types.Overview `json:"OverviewState"`
	types.OverviewOpenedOrClosed
}

//go:generate ./generate/bin/request casts Casts []types.Cast

// Documented, but doens't work.
// TODO
type castsResponse struct {
	Casts []types.Cast `json:"Casts"`
}

type simpleAction struct {
	b []byte
}

func newSimpleAction(b []byte) simpleAction {
	var a simpleAction
	a.b = b
	return a
}

func (a simpleAction) MarshalJSON() ([]byte, error) {
	return a.b, nil // Assume it wont be mutated by encoding/json
}

func (a simpleAction) response() Response { return new(string) }

type action struct {
	j struct {
		Action any `json:"Action"`
	}
}

func newAction(i any) action {
	var a action
	a.j.Action = i
	return a
}

func (a action) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.j)
}

func (a action) response() Response { return new(string) }

//go:generate go build -o generate/bin/action ./generate/action
//go:generate ./generate/bin/action create

//go:generate ./generate/bin/action PowerOffMonitors
//go:generate ./generate/bin/action PowerOnMonitors
//go:generate ./generate/bin/action ToggleKeyboardShortcutsInhibit
//go:generate ./generate/bin/action FocusWindowPrevious
//go:generate ./generate/bin/action FocusColumnLeft
//go:generate ./generate/bin/action FocusColumnRight
//go:generate ./generate/bin/action FocusColumnFirst
//go:generate ./generate/bin/action FocusColumnLast
//go:generate ./generate/bin/action FocusColumnRightOrFirst
//go:generate ./generate/bin/action FocusColumnLeftOrLast
//go:generate ./generate/bin/action FocusWindowOrMonitorUp
//go:generate ./generate/bin/action FocusWindowOrMonitorDown
//go:generate ./generate/bin/action FocusColumnOrMonitorLeft
//go:generate ./generate/bin/action FocusColumnOrMonitorRight
//go:generate ./generate/bin/action FocusWindowDown
//go:generate ./generate/bin/action FocusWindowUp
//go:generate ./generate/bin/action FocusWindowDownOrColumnLeft
//go:generate ./generate/bin/action FocusWindowDownOrColumnRight
//go:generate ./generate/bin/action FocusWindowUpOrColumnLeft
//go:generate ./generate/bin/action FocusWindowUpOrColumnRight
//go:generate ./generate/bin/action FocusWindowOrWorkspaceDown
//go:generate ./generate/bin/action FocusWindowOrWorkspaceUp
//go:generate ./generate/bin/action FocusWindowTop
//go:generate ./generate/bin/action FocusWindowBottom
//go:generate ./generate/bin/action FocusWindowDownOrTop
//go:generate ./generate/bin/action FocusWindowUpOrBottom
//go:generate ./generate/bin/action MoveColumnLeft
//go:generate ./generate/bin/action MoveColumnRight
//go:generate ./generate/bin/action MoveColumnToFirst
//go:generate ./generate/bin/action MoveColumnToLast
//go:generate ./generate/bin/action MoveColumnLeftOrToMonitorLeft
//go:generate ./generate/bin/action MoveColumnRightOrToMonitorRight
//go:generate ./generate/bin/action MoveWindowDown
//go:generate ./generate/bin/action MoveWindowUp
//go:generate ./generate/bin/action MoveWindowDownOrToWorkspaceDown
//go:generate ./generate/bin/action MoveWindowUpOrToWorkspaceUp
//go:generate ./generate/bin/action ConsumeWindowIntoColumn
//go:generate ./generate/bin/action ExpelWindowFromColumn
//go:generate ./generate/bin/action SwapWindowRight
//go:generate ./generate/bin/action SwapWindowLeft
//go:generate ./generate/bin/action ToggleColumnTabbedDisplay
//go:generate ./generate/bin/action CenterColumn
//go:generate ./generate/bin/action CenterVisibleColumns
//go:generate ./generate/bin/action FocusWorkspaceDown
//go:generate ./generate/bin/action FocusWorkspaceUp
//go:generate ./generate/bin/action FocusWorkspacePrevious
//go:generate ./generate/bin/action MoveWorkspaceDown
//go:generate ./generate/bin/action MoveWorkspaceUp
//go:generate ./generate/bin/action FocusMonitorLeft
//go:generate ./generate/bin/action FocusMonitorRight
//go:generate ./generate/bin/action FocusMonitorDown
//go:generate ./generate/bin/action FocusMonitorUp
//go:generate ./generate/bin/action FocusMonitorPrevious
//go:generate ./generate/bin/action FocusMonitorNext
//go:generate ./generate/bin/action MoveWindowToMonitorLeft
//go:generate ./generate/bin/action MoveWindowToMonitorRight
//go:generate ./generate/bin/action MoveWindowToMonitorDown
//go:generate ./generate/bin/action MoveWindowToMonitorUp
//go:generate ./generate/bin/action MoveWindowToMonitorPrevious
//go:generate ./generate/bin/action MoveWindowToMonitorNext
//go:generate ./generate/bin/action MoveColumnToMonitorLeft
//go:generate ./generate/bin/action MoveColumnToMonitorRight
//go:generate ./generate/bin/action MoveColumnToMonitorDown
//go:generate ./generate/bin/action MoveColumnToMonitorUp
//go:generate ./generate/bin/action MoveColumnToMonitorPrevious
//go:generate ./generate/bin/action MoveColumnToMonitorNext
//go:generate ./generate/bin/action SwitchPresetColumnWidth
//go:generate ./generate/bin/action SwitchPresetColumnWidthBack
//go:generate ./generate/bin/action MaximizeColumn
//go:generate ./generate/bin/action ExpandColumnToAvailableWidth
//go:generate ./generate/bin/action ShowHotkeyOverlay
//go:generate ./generate/bin/action MoveWorkspaceToMonitorLeft
//go:generate ./generate/bin/action MoveWorkspaceToMonitorRight
//go:generate ./generate/bin/action MoveWorkspaceToMonitorDown
//go:generate ./generate/bin/action MoveWorkspaceToMonitorUp
//go:generate ./generate/bin/action MoveWorkspaceToMonitorPrevious
//go:generate ./generate/bin/action MoveWorkspaceToMonitorNext
//go:generate ./generate/bin/action ToggleDebugTint
//go:generate ./generate/bin/action DebugToggleOpaqueRegions
//go:generate ./generate/bin/action DebugToggleDamage
//go:generate ./generate/bin/action FocusFloating
//go:generate ./generate/bin/action FocusTiling
//go:generate ./generate/bin/action SwitchFocusBetweenFloatingAndTiling
//go:generate ./generate/bin/action ClearDynamicCastTarget
//go:generate ./generate/bin/action ToggleOverview
//go:generate ./generate/bin/action OpenOverview
//go:generate ./generate/bin/action CloseOverview
//go:generate ./generate/bin/action LoadConfigFile

//go:generate ./generate/bin/action Quit skip_confirmation:skipConfirm:bool
//go:generate ./generate/bin/action Spawn command:[]string
//go:generate ./generate/bin/action SpawnSh command:string

func ActionDoScreenTransition(delay *time.Duration) Request {
	if delay == nil {
		return newSimpleAction([]byte(`{"Action":{"DoScreenTransition":{}}}`))
	}
	ms := uint16(delay.Milliseconds())
	return newAction(map[string]map[string]uint16{
		"DoScreenTransition": {"delay_ms": ms},
	})
}

//go:generate ./generate/bin/action Screenshot show_pointer:bool path:*string
//go:generate ./generate/bin/action ScreenshotScreen write_to_disk:write:bool show_pointer:showPointer:bool path:*string
//go:generate ./generate/bin/action ScreenshotWindow id:*u64 write_to_disk:write:bool show_pointer:showPointer:bool path:*string
//go:generate ./generate/bin/action CloseWindow id:*u64
//go:generate ./generate/bin/action FullscreenWindow id:*u64
//go:generate ./generate/bin/action ToggleWindowedFullscreen id:*u64
//go:generate ./generate/bin/action FocusWindow id:u64
//go:generate ./generate/bin/action FocusWindowInColumn index:u8
//go:generate ./generate/bin/action FocusColumn index:usize
//go:generate ./generate/bin/action MoveColumnToIndex index:usize
//go:generate ./generate/bin/action ConsumeOrExpelWindowLeft id:*u64
//go:generate ./generate/bin/action ConsumeOrExpelWindowRight id:*u64
//go:generate ./generate/bin/action SetColumnDisplay display:types.ColumnDisplay
//go:generate ./generate/bin/action CenterWindow id:*u64
//go:generate ./generate/bin/action FocusWorkspace reference:workspaceRef:types.WorkspaceRef
//go:generate ./generate/bin/action MoveWindowToWorkspaceDown focus:bool
//go:generate ./generate/bin/action MoveWindowToWorkspaceUp focus:bool
//go:generate ./generate/bin/action MoveWindowToWorkspace id:*u64 reference:workspaceRef:types.WorkspaceRef focus:bool
//go:generate ./generate/bin/action MoveColumnToWorkspaceDown focus:bool
//go:generate ./generate/bin/action MoveColumnToWorkspaceUp focus:bool
//go:generate ./generate/bin/action MoveColumnToWorkspace reference:workspaceRef:types.WorkspaceRef focus:bool
//go:generate ./generate/bin/action MoveWorkspaceToIndex index:usize reference:workspaceRef:types.WorkspaceRef
//go:generate ./generate/bin/action SetWorkspaceName name:string workspace:workspaceRef:*types.WorkspaceRef
//go:generate ./generate/bin/action UnsetWorkspaceName reference:workspaceRef:*types.WorkspaceRef
//go:generate ./generate/bin/action FocusMonitor output:string
//go:generate ./generate/bin/action MoveWindowToMonitor id:*u64 output:string
//go:generate ./generate/bin/action MoveColumnToMonitor output:string
//go:generate ./generate/bin/action SetWindowWidth id:*u64 change:types.SizeChange
//go:generate ./generate/bin/action SetWindowHeight id:*u64 change:types.SizeChange
//go:generate ./generate/bin/action ResetWindowHeight id:*u64
//go:generate ./generate/bin/action SwitchPresetWindowWidth id:*u64
//go:generate ./generate/bin/action SwitchPresetWindowWidthBack id:*u64
//go:generate ./generate/bin/action SwitchPresetWindowHeight id:*u64
//go:generate ./generate/bin/action SwitchPresetWindowHeightBack id:*u64
//go:generate ./generate/bin/action MaximizeWindowToEdges id:*u64
//go:generate ./generate/bin/action SetColumnWidth change:types.SizeChange
//go:generate ./generate/bin/action SwitchLayout layout:types.LayoutSwitchTarget
//go:generate ./generate/bin/action MoveWorkspaceToMonitor output:string reference:workspaceRef:*types.WorkspaceRef
//go:generate ./generate/bin/action ToggleWindowFloating id:*u64
//go:generate ./generate/bin/action MoveWindowToFloating id:*u64
//go:generate ./generate/bin/action MoveWindowToTiling id:*u64
//go:generate ./generate/bin/action MoveFloatingWindow id:*u64 x: y:types.PositionChange
//go:generate ./generate/bin/action ToggleWindowRuleOpacity id:*u64
//go:generate ./generate/bin/action SetDynamicCastWindow id:*u64
//go:generate ./generate/bin/action SetDynamicCastMonitor output:*string
//go:generate ./generate/bin/action StopCast session_id:sessionID:u64
//go:generate ./generate/bin/action ToggleWindowUrgent id:u64
//go:generate ./generate/bin/action SetWindowUrgent id:u64
//go:generate ./generate/bin/action UnsetWindowUrgent id:u64

type outputAction struct {
	r *outputConfigChangedResponse
	j struct {
		Output struct {
			Output string `json:"output"`
			Action any    `json:"action"`
		} `json:"Output"`
	}
}

func newOutputAction(o string, i any) outputAction {
	var a outputAction
	a.r = new(outputConfigChangedResponse)
	a.j.Output.Output = o
	a.j.Output.Action = i
	return a
}

func (a outputAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.j)
}

func (a outputAction) response() Response { return a.r }

//go:generate go build -o generate/bin/output-action ./generate/output-action
//go:generate ./generate/bin/output-action create

//go:generate ./generate/bin/output-action Off
//go:generate ./generate/bin/output-action On
//go:generate ./generate/bin/output-action Mode mode:types.ModeToSet
//go:generate ./generate/bin/output-action CustomMode mode:types.ConfiguredMode
//go:generate ./generate/bin/output-action Modeline clock:f64 hdisplay:u16 hsync_start:u16 hsync_end:u16 htotal:u16 vdisplay:u16 vsync_start:u16 vsync_end:u16 vtotal:u16 hsync_polarity:types.HSyncPolarity vsync_polarity:types.VSyncPolarity
//go:generate ./generate/bin/output-action Scale scale:types.ScaleToSet
//go:generate ./generate/bin/output-action Transform transform:types.Transform
//go:generate ./generate/bin/output-action Position position:types.PositionToSet
//go:generate ./generate/bin/output-action Vrr:VRR vrr:types.VRRToSet
