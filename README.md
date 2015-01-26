Windstorm
=========
Windstorm is a simple window creation library for Go. It creates OpenGL
contexts and handles user input.

Why?
----
Usually, Go programs use GLFW bindings to create windows and get user input.
This approach can work very well, and there are some great bindings out there
like [this one](http://github.com/go-gl/glfw3) that make it a fairly painless
process. However, bindings, by necessity, don't fit in with the Go way of doing
things. Bindings are not Go-gettable because they require C libraries to be
installed first and their design generally leans towards a more C-like
approach. Windstorm still requires cgo because it needs to interact with the
libraries that the operating system supplies for window management, but it's
still one less level of abstraction between Go programs and creating windows.

Windstorm hopes to offer an alternative to bindings that is Go-gettable (for
the most part, more details below) and more similar to the way Go programmers
expect to interact with libraries. It will be a while before Windstorm reaches
GLFW's level of feature-completeness and stability, but I believe it's well
worth the effort.

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

Gopher Gala Info
----------------
Full disclosure: Work on this project is based on an earlier attempt at the
idea, but I'm doing everything over again, so I believe I'm in the green
rule-wise. I'm also designing the library very differently both in code and
library structure. This does give me a bit of a head-start though, as I've
already done a lot of the research.
