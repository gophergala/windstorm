Windstorm
=========
Windstorm is a simple window creation library for Go. It creates OpenGL
contexts and handles user input.

Why?
----
When I created this, there weren't as many great Go solutions to this problem.
If you need a window creation library, you should use
[the shiny library](http://golang.org/x/exp/shiny), which is actively
developed by some big names in the Go community.

Installation
------------
In many cases, Windstorm can be installed with a simple call to "go get".
However, some Linux distributions may not come with the X development
libraries. This should be an easy install from your package manager of choice.

Use
---
In many ways, Windstorm is used in a similar way to other, similar libraries.
However, when it comes to user input, it takes a very different approach.
Instead of using callbacks, Windstorm uses channels to communicate events. This
allows for greater control over how your program is structured. Input could be
managed comfortably within the main loop, handled in a separate goroutine, or
set up to simply run functions in a callback-like manner, if you prefer.

As far as drawing is concerned, you're given a OpenGL context to do anything
you'd like with. Just make sure to run SwapBuffers to let Windstorm know you're
done drawing a frame.

TODO
----
 - Windows support (In progress)
 - Mac support
 - Joystick input
 - Fullscreen support
