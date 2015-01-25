// simple.go
// A simple program that draws a red box on screen, has it move around relative
// to the position of the mouse, and quits when the user hits escape.
package main

import (
	"fmt"
	gl "github.com/chsc/gogl/gl21"
	"github.com/gophergala/windstorm"
)

func drawSquare(x, y int) {

	normX := ((gl.Float(x)) / 640.0)
	normY := ((gl.Float(y)) / 480.0)

	gl.Color3f(1.0, 0.0, 0.0)
	gl.Begin(gl.LINE_LOOP)
	gl.Vertex2f(normX-0.05, normY-0.05)
	gl.Vertex2f(normX+0.05, normY-0.05)
	gl.Vertex2f(normX+0.05, normY+0.05)
	gl.Vertex2f(normX-0.05, normY+0.05)
	gl.End()
}

func main() {

	window, err := windstorm.NewWindow(400, 400, "Simple Windstorm Test")
	if err != nil {
		fmt.Println(err)
		return
	}

	window.Show()

	if err = window.CreateContext(); err != nil {
		fmt.Println(err)
		return
	}
	if err = window.MakeContextCurrent(); err != nil {
		fmt.Println(err)
		return
	}

	window.SetRecievesEvents(true)

	if err = gl.Init(); err != nil {
		fmt.Println(err)
		return
	}

	var squareX, squareY int

	gl.ClearColor(1.0, 1.0, 1.0, 1.0)

	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		drawSquare(squareX, squareY)

		if err = window.SwapBuffers(); err != nil {
			fmt.Println(err)
			return
		}
		window.UpdateEvents()

		for stop := false; !stop; {
			select {
			case <-window.OnClose:
				window.Close()
				windstorm.Stop()
				return
			case event := <-window.OnKeyboard:
				if event.Key == windstorm.KeyEscape {
					fmt.Println("quitting")
					window.Close()
					windstorm.Stop()
					return
				}
			case event := <-window.OnMouseMove:
				squareX = event.X
				squareY = event.Y
			default:
				stop = true
			}
		}
	}

	windstorm.Stop()
}
