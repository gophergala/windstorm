package windstorm

import (
	"errors"
)

type Window struct {
	width, height int
	title         string
	cWin          cWindow
	cCont         cContext
}

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

func (window *Window) UpdateEvents() error {

	err := cUpdateEvents(window.cWin)

	return err
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
