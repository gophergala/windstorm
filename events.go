package windstorm

// CloseEvent represents a request for a window to close.
type CloseEvent struct {
}

// ResizeEvent represents a user or window manager triggered resizing of a
// window.
type ResizeEvent struct {
	Width  int
	Height int
}

// KeyboardEvent represents a user triggered keyboard press or release event.
type KeyboardEvent struct {
	Key    Key
	Action Action
}

// MouseMoveEvent represents a movement of the mouse cursor.
type MouseMoveEvent struct {
	X int
	Y int
}

// MouseButtonEvent represents a user triggered mouse button press or release
// event.
type MouseButtonEvent struct {
	Button MouseButton
	Action Action
}

// FocusEvent represents a change in window focus.
type FocusEvent struct {
	Focused bool
}

// MouseLeaveWindowEvent represents an event triggered by the mouse leaving a
// window.
type MouseLeaveWindowEvent struct {
}

// MouseEnterWindowEvent represents an event triggered by the mouse entering
// back into a window.
type MouseEnterWindowEvent struct {
	X int
	Y int
}
