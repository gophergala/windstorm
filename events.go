package windstorm

type ResizeEvent struct {
	Width  int
	Height int
}

type KeyboardEvent struct {
	Key    Key
	Action Action
}

type MouseMoveEvent struct {
	X int
	Y int
}

type MouseButtonEvent struct {
	Button MouseButton
	Action Action
}

type FocusEvent struct {
	Focused bool
}
