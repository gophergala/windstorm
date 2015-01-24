package windstorm

import (
	"errors"
)

type Window struct {
	width, height int
	title         string
}

func NewWindow(width, height int, title string) (Window, error) {

	var window Window

	if width <= 0 || height <= 0 {
		return Window{}, errors.New("invalid window size attributes")
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
