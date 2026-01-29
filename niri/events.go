package niri

import "time"

type EventHandler interface {
	Workspaces([]Workspace)
	WorkspaceUrgency(wsid WorkspaceID, urgent bool)
	WorkspaceFocus(wsid WorkspaceID, focus bool)
	WorkspaceActiveWindow(wsid WorkspaceID, wid WindowID)
	Windows([]Window)
	Window(Window)
	WindowClosed(WindowID)
	WindowFocus(WindowID)
	WindowFocusTimestamp(wid WindowID, epoch time.Duration)
	WindowUrgency(wid WindowID, urgent bool)
	WindowLayouts(map[WindowID]WindowLayout)
	KeyboardLayouts(index int, layouts []string)
	KeyboardLayoutSwitched(index int)
	Overview(open bool)
	ConfigLoaded(ok bool)
	Screenshot(path string)
	// Casts(*CastsChanged)
	// Cast(*CastStartedOrChanged)
	// CastStopped(*CastStopped)
}

type NoOpEventHandler struct{}

func (NoOpEventHandler) Workspaces([]Workspace)                       {}
func (NoOpEventHandler) WorkspaceUrgency(WorkspaceID, bool)           {}
func (NoOpEventHandler) WorkspaceFocus(WorkspaceID, bool)             {}
func (NoOpEventHandler) WorkspaceActiveWindow(WorkspaceID, WindowID)  {}
func (NoOpEventHandler) Windows([]Window)                             {}
func (NoOpEventHandler) Window(Window)                                {}
func (NoOpEventHandler) WindowClosed(WindowID)                        {}
func (NoOpEventHandler) WindowFocus(WindowID)                         {}
func (NoOpEventHandler) WindowFocusTimestamp(WindowID, time.Duration) {}
func (NoOpEventHandler) WindowUrgency(WindowID, bool)                 {}
func (NoOpEventHandler) WindowLayouts(map[WindowID]WindowLayout)      {}
func (NoOpEventHandler) KeyboardLayouts(int, []string)                {}
func (NoOpEventHandler) KeyboardLayoutSwitched(int)                   {}
func (NoOpEventHandler) Overview(bool)                                {}
func (NoOpEventHandler) ConfigLoaded(bool)                            {}
func (NoOpEventHandler) Screenshot(string)                            {}

// func (NoOpEventHandler) Casts(*CastsChanged)                                 {}
// func (NoOpEventHandler) CastStartedOr(*CastStartedOrChanged)                 {}
// func (NoOpEventHandler) CastStopped(*CastStopped)                            {}

var _ EventHandler = NoOpEventHandler{}
