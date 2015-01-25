package windstorm

import (
	"errors"
)

type Window struct {
	width, height int
	title         string
	cWin          cWindow
	cCont         cContext

	keyStates         map[Key]Action
	mouseX, mouseY    int
	mouseButtonStates map[MouseButton]Action
	focused           bool
	shouldClose       bool

	OnClose       chan CloseEvent
	OnResize      chan ResizeEvent
	OnKeyboard    chan KeyboardEvent
	OnMouseMove   chan MouseMoveEvent
	OnMouseButton chan MouseButtonEvent
	OnFocus       chan FocusEvent
}

var windows map[cWindow]*Window

func NewWindow(width, height int, title string) (Window, error) {

	var window Window
	var err error

	if width <= 0 || height <= 0 {
		return Window{}, errors.New("invalid window size attributes")
	}

	window.cWin, err = cNewWindow(width, height, title)
	if err != nil {
		return Window{}, errors.New("unable to create window")
	}

	window.width = width
	window.height = height
	window.title = title

	window.keyStates = make(map[Key]Action)
	window.mouseButtonStates = make(map[MouseButton]Action)

	window.OnClose = make(chan CloseEvent, 256)
	window.OnResize = make(chan ResizeEvent, 256)
	window.OnKeyboard = make(chan KeyboardEvent, 256)
	window.OnMouseMove = make(chan MouseMoveEvent, 256)
	window.OnMouseButton = make(chan MouseButtonEvent, 256)
	window.OnFocus = make(chan FocusEvent, 256)

	return window, nil
}

func (window *Window) Width() int {

	return window.width
}

func (window *Window) Height() int {

	return window.height
}

func (window *Window) Title() string {

	return window.title
}

func (window *Window) Show() error {

	err := cShowWindow(window.cWin)

	return err
}

func (window *Window) Hide() error {

	err := cHideWindow(window.cWin)

	return err
}

func (window *Window) Close() error {

	err := cCloseWindow(window.cWin)

	return err
}

func (window *Window) SetRecievesEvents(recieves bool) {

	if recieves {
		windows[window.cWin] = window
	} else {
		delete(windows, window.cWin)
	}
}

func (window *Window) SetTitle(title string) error {

	return cSetWindowTitle(title, window.cWin)
}

func (window *Window) SetSize(width, height int) error {

	return cResizeWindow(width, height, window.cWin)
}

func (window *Window) UpdateEvents() error {

	for stop := false; !stop; {
		select {
		case <-window.OnResize:
		case <-window.OnKeyboard:
		default:
			stop = true
		}
	}

	err := cUpdateEvents(window.cWin)

	return err
}

func (window *Window) KeyState(key Key) Action {

	return window.keyStates[key]
}

func (window *Window) MousePosition() (int, int) {

	return window.mouseX, window.mouseY
}

func (window *Window) MouseButtonState(button MouseButton) Action {

	return window.mouseButtonStates[button]
}

func (window *Window) InFocus() bool {

	return window.focused
}

func (window *Window) ShouldClose() bool {

	return window.shouldClose
}

func (window *Window) CreateContext() error {

	var err error
	window.cCont, err = cCreateContext()

	return err
}

func (window *Window) MakeContextCurrent() error {

	err := cMakeContextCurrent(window.cWin, window.cCont)

	return err
}

func (window *Window) SwapBuffers() error {

	err := cSwapBuffers(window.cWin)

	return err
}

func (window *Window) onClose() {

	window.shouldClose = true

	select {
	case window.OnClose <- CloseEvent{}:
	default:
	}
}

func (window *Window) onResize(width, height int) {

	window.width = width
	window.height = height

	select {
	case window.OnResize <- ResizeEvent{Width: width, Height: height}:
	default:
	}
}

func (window *Window) onKeyboard(key Key, action Action) {

	window.keyStates[key] = action

	select {
	case window.OnKeyboard <- KeyboardEvent{Key: key, Action: action}:
	default:
	}
}

func (window *Window) onMouseMove(x, y int) {

	window.mouseX = x
	window.mouseY = y

	select {
	case window.OnMouseMove <- MouseMoveEvent{X: x, Y: y}:
	default:
	}
}

func (window *Window) onMouseButton(button MouseButton, action Action) {

	window.mouseButtonStates[button] = action

	select {
	case window.OnMouseButton <- MouseButtonEvent{Button: button, Action: action}:
	default:
	}
}

func (window *Window) onFocus(focused bool) {

	window.focused = focused

	select {
	case window.OnFocus <- FocusEvent{Focused: focused}:
	default:
	}
}
