#include <X11/Xlib.h>
#include <errno.h>
#include <GL/glx.h>
#include <stdio.h>
#include "_cgo_export.h"

Display *display;
Window rootWindow;
int screen;
long eventMask = KeyPressMask | KeyReleaseMask | PointerMotionMask | FocusChangeMask;
GLint glAttribs[] = { GLX_RGBA, GLX_DEPTH_SIZE, 24, GLX_DOUBLEBUFFER, None };
XSetWindowAttributes winAttribs;
XVisualInfo *vi;
char *errorMsg;

typedef Window WindstormWindow;
typedef GLXContext WindstormContext;

const int Press = 1;
const int Release = 2;

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

void WindstormSetWindowTitle(char *title, WindstormWindow window) {

	XStoreName(display, window, title);

	errno = 0;
}

int predicateFunc(Display* d, XEvent* e, XPointer p) {

	return True;
}

Bool peekEvents(XEvent *event) {

	if(XPending(display) > 0) {
		XPeekEvent(display, event);
		return True;
	}

	return False;
}

void WindstormUpdateEvents(WindstormWindow window) {

	XWindowAttributes attribs;
	XGetWindowAttributes(display, window, &attribs);

	XEvent event;

	int lastKeyReleased = -1;
	int keyVal;
	while(XCheckIfEvent(display, &event, predicateFunc, NULL)) {
		switch(event.type) {
		case ClientMessage:
			XDestroyWindow(display, window);
			lastKeyReleased = -1;
			break;
		case KeyPress:
			// The lastKeyReleased value will be the value of the last key
			// released if a KeyRelease event was the last thing to happen.
			// If that value is the same as the value of the KeyPress event,
			// the KeyPress event will be ignored to prevent the behavior
			// described above.
			keyVal = XLookupKeysym(&event.xkey, 0);
			if(keyVal != lastKeyReleased) {
				keyboardEvent(keyVal, Press, window);
			}
			break;
		case KeyRelease:
			// X11 sends "key held down" events in the form of a key release
			// event and a key press event following directly after. In
			// order to avoid this, the next event has to be reviewed to see if
			// it follows this pattern.
			keyVal = XLookupKeysym(&event.xkey, 0);
			XEvent nextEvent;
			if(peekEvents(&nextEvent) == True && nextEvent.type == KeyPress && XLookupKeysym(&nextEvent.xkey, 0) == keyVal) {
				//keyboardEvent(keyVal, Hold, window);
			} else {
				keyboardEvent(keyVal, Release, window);
			}
			lastKeyReleased = keyVal;
			break;
		case MotionNotify:
			mouseMoveEvent(event.xmotion.x, attribs.height - event.xmotion.y, window);
			break;
		case FocusIn:
			focusEvent(1, window);
			break;
		case FocusOut:
			focusEvent(0, window);
			break;
		}
	}

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
