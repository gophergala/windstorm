#include <X11/Xlib.h>
#include <errno.h>
#include <GL/glx.h>
#include <stdio.h>
#include "_cgo_export.h"

Display *display;
Window rootWindow;
int screen;
long eventMask = ExposureMask;
GLint glAttribs[] = { GLX_RGBA, GLX_DEPTH_SIZE, 24, GLX_DOUBLEBUFFER, None };
XSetWindowAttributes winAttribs;
XVisualInfo *vi;
char *errorMsg;

typedef Window WindstormWindow;
typedef GLXContext WindstormContext;

int errorHandler(Display *display, XErrorEvent *event) {

	char *message = "                                 ";
	XGetErrorText(display, event->error_code, message, 32);
	printf(message);
}

void WindstormInit() {

	display = XOpenDisplay(NULL);
	if(display == NULL) {
		errorMsg = "could not connect to X server";
		errno = -1;
		return;
	}

	XSetErrorHandler(errorHandler);

	screen = DefaultScreen(display);
	rootWindow = DefaultRootWindow(display);

	vi = glXChooseVisual(display, 0, glAttribs);
	if(vi == NULL) {
		errorMsg = "could not find proper GLX visual";
		errno = -1;
		return;
	}

	Colormap colorMap = XCreateColormap(display, rootWindow, vi->visual, AllocNone);
	winAttribs.colormap = colorMap;
	winAttribs.event_mask = eventMask;

	errno = 0;
}

WindstormWindow WindstormNewWindow(int width, int height, char *title) {

	Window window = XCreateWindow(display, rootWindow, 0, 0, width, height, 0,
		vi->depth, InputOutput, vi->visual, CWColormap | CWEventMask, &winAttribs);

	XSetStandardProperties(display, window, title, title, None, NULL, 0, NULL);

	errno = 0;

	return window;
}

void WindstormShowWindow(WindstormWindow window) {

	XMapWindow(display, window);

	errno = 0;
}

void WindstormHideWindow(WindstormWindow window) {

	XUnmapWindow(display, window);

	errno = 0;
}

int predicateFunc(Display* d, XEvent* e, XPointer p) {

	return True;
}

void WindstormUpdateEvents(WindstormWindow window) {

	XEvent event;

	while(XCheckIfEvent(display, &event, predicateFunc, NULL)) {
		switch(event.type) {
		case ClientMessage:
			XDestroyWindow(display, window);
			break;
		}
	}

	XWindowAttributes attribs;
	XGetWindowAttributes(display, window, &attribs);

	// Instead of using a ResizeRedirectMask to check if the window size has
	// changed, this check is done manually. Using ResizeRedirectMask keeps the
	// default behaviors from happening, which would then have to be emulated.
	// TODO: Find a better way to do this.
	resizeEvent(attribs.width, attribs.height, window);

	errno = 0;

	return;
}

void WindstormCloseWindow(WindstormWindow window) {

	XDestroyWindow(display, window);

	errno = 0;
}

WindstormContext WindstormCreateContext() {

	GLXContext context = glXCreateContext(display, vi, NULL, True);
	if(context == NULL) {
		errorMsg = "could not create OpenGL context";
		errno = -1;
		return NULL;
	}

	errno = 0;
	return context;
}

void WindstormMakeContextCurrent(WindstormWindow window, WindstormContext context) {

	if(glXMakeCurrent(display, window, context) == False) {
		errorMsg = "could not make OpenGL context current";
		errno = -1;
		return;
	}

	errno = 0;
}

void WindstormSwapBuffers(WindstormWindow window) {

	glXSwapBuffers(display, window);

	errno = 0;
}

void WindstormStop() {

	XCloseDisplay(display);
	errno = 0;
}
