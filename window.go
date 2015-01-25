package windstorm

import (
	"errors"
)

// Window is an object representing a window with an OpenGL context inside of
// it. Most operations are done through a Window object. OnX channels are
// exposed through this object that may be read for user input. These channels
// should never be written to.
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
	mouseInWindow     bool

	OnClose            chan CloseEvent
	OnResize           chan ResizeEvent
	OnKeyboard         chan KeyboardEvent
	OnMouseMove        chan MouseMoveEvent
	OnMouseButton      chan MouseButtonEvent
	OnFocus            chan FocusEvent
	OnMouseEnterWindow chan MouseEnterWindowEvent
	OnMouseLeaveWindow chan MouseLeaveWindowEvent
}

var windows map[cWindow]*Window

// NewWindow creates a new Window object with the given width, height, and
// title. Width and height attributes must be above zero. The window will not
// immediately appear on screen. The Show method must be called first. The
// window will also not capture user input by default. The SetRetrievesInput
// method must be called first.
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
	window.OnMouseEnterWindow = make(chan MouseEnterWindowEvent, 256)
	window.OnMouseLeaveWindow = make(chan MouseLeaveWindowEvent, 256)

	return window, nil
}

// Width returns the current width of the window in pixels.
func (window *Window) Width() int {

	return window.width
}

// Height returns the current height of the window in pixels.
func (window *Window) Height() int {

	return window.height
}

// Title returns the current title of the window.
func (window *Window) Title() string {

	return window.title
}

// MouseInWindow returns true if the cursor currently resides in the window
// space.
func (window *Window) MouseInWindow() bool {

	return window.mouseInWindow
}

// Show displays the window on screen.
func (window *Window) Show() error {

	err := cShowWindow(window.cWin)

	return err
}

// Hide stops the window from being displayed on screen, but does not close it.
func (window *Window) Hide() error {

	err := cHideWindow(window.cWin)

	return err
}

// Close closes the window. Drawing should not be performed on it's context
// after this method is called. User input will stop being fed to the window,
// as well.
func (window *Window) Close() error {

	err := cCloseWindow(window.cWin)

	return err
}

// SetRecievesEvents toggles whether the window should recieve user input
// events or not. The default value is false.
func (window *Window) SetRecievesEvents(recieves bool) {

	if recieves {
		windows[window.cWin] = window
	} else {
		delete(windows, window.cWin)
	}
}

// SetTitle sets the title of the window.
func (window *Window) SetTitle(title string) error {

	return cSetWindowTitle(title, window.cWin)
}

// SetSize sets the width and height of the window in pixels. Valid values
// must be above zero.
func (window *Window) SetSize(width, height int) error {

	return cResizeWindow(width, height, window.cWin)
}

// UpdateEvents checks for new user input events.
func (window *Window) UpdateEvents() error {

	for stop := false; !stop; {
		select {
		case <-window.OnClose:
		case <-window.OnResize:
		case <-window.OnKeyboard:
		case <-window.OnMouseMove:
		case <-window.OnMouseButton:
		case <-window.OnFocus:
		default:
			stop = true
		}
	}

	err := cUpdateEvents(window.cWin)

	return err
}

// KeyState returns the current state of a Key.
func (window *Window) KeyState(key Key) Action {

	return window.keyStates[key]
}

// MousePosition returns the position of the cursor.
func (window *Window) MousePosition() (int, int) {

	return window.mouseX, window.mouseY
}

// MouseButtonState returns the state of a mouse button.
func (window *Window) MouseButtonState(button MouseButton) Action {

	return window.mouseButtonStates[button]
}

// InFocus returns true if the window is currently in focus.
func (window *Window) InFocus() bool {

	return window.focused
}

// ShouldClose returns true if the window has recieved a close request from
// the user or the OS.
func (window *Window) ShouldClose() bool {

	return window.shouldClose
}

// CreateContext creates a new OpenGL context.
func (window *Window) CreateContext() error {

	var err error
	window.cCont, err = cCreateContext()

	return err
}

// MakeContextCurrent makes the window's context current. Future OpenGL draw
// operations will be done on this context.
func (window *Window) MakeContextCurrent() error {

	err := cMakeContextCurrent(window.cWin, window.cCont)

	return err
}

// SwapBuffers updates the display in the window.
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

	window.mouseInWindow = true

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

func (window *Window) onMouseEnterWindow(x, y int) {

	window.mouseInWindow = true

	select {
	case window.OnMouseEnterWindow <- MouseEnterWindowEvent{X: x, Y: y}:
	default:
	}
}

func (window *Window) onMouseLeaveWindow() {

	window.mouseInWindow = false

	select {
	case window.OnMouseLeaveWindow <- MouseLeaveWindowEvent{}:
	default:
	}
}
