package windstorm

type ResizeEvent struct {
	Width  int
	Height int
}

type KeyboardEvent struct {
	Key    int
	Action Action
}