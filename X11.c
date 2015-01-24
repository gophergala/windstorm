#include <X11/Xlib.h>
#include <errno.h>
#include <GL/glx.h>
#include <stdio.h>
#include "_cgo_export.h"

Display *display;
Window rootWindow;
int screen;
GLint glAttribs[] = { GLX_RGBA, GLX_DEPTH_SIZE, 24, GLX_DOUBLEBUFFER, None };
XSetWindowAttributes winAttribs;
XVisualInfo *vi;

typedef Window WindstormWindow;

int errorHandler(Display *display, XErrorEvent *event) {

	char *message = "                                 ";
	XGetErrorText(display, event->error_code, message, 32);
	printf(message);
}

void WindstormInit() {

	display = XOpenDisplay(NULL);
	if(display == NULL) {
		errno = -1;
		return;
	}

	XSetErrorHandler(errorHandler);

	screen = DefaultScreen(display);
	rootWindow = DefaultRootWindow(display);

	vi = glXChooseVisual(display, 0, glAttribs);
	if(vi == NULL) {
		errno = -1;
		return;
	}

	Colormap colorMap = XCreateColormap(display, rootWindow, vi->visual, AllocNone);
	winAttribs.colormap = colorMap;

	errno = 0;
}

WindstormWindow WindstormNewWindow(int width, int height, char *title) {

	Window window = XCreateWindow(display, rootWindow, 0, 0, width, height, 0,
		vi->depth, InputOutput, vi->visual, CWColormap, &winAttribs);

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

int WindstormUpdateEvents(WindstormWindow window) {

	XEvent event;

	while(XCheckIfEvent(display, &event, predicateFunc, NULL)) {
		switch(event.type) {
		case ClientMessage:
			XDestroyWindow(display, window);
			break;
		}
	}

	errno = 0;

	return 0;
}

void WindstormCloseWindow(WindstormWindow window) {

	XDestroyWindow(display, window);

	errno = 0;
}
