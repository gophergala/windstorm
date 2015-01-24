#include <X11/Xlib.h>
#include <errno.h>

Display *display;
Window rootWindow;
int screen;

void WindstormInit() {

	display = XOpenDisplay(NULL);
	if(display == NULL) {
		errno = -1;
		return;
	}

	screen = DefaultScreen(display);
	rootWindow = DefaultRootWindow(display);

	errno = 0;
}
