package windstorm

/*
#include <X11/Xlib.h>
#include <GL/glx.h>
typedef Window WindstormWindow;
typedef GLXContext WindstormContext;

extern void WindstormInit();
extern WindstormWindow WindstormNewWindow(int, int, char*);
extern void WindstormShowWindow(WindstormWindow);
extern void WindstormHideWindow(WindstormWindow);
extern void WindstormUpdateEvents(WindstormWindow);
extern void WindstormCloseWindow(WindstormWindow);
extern WindstormContext WindstormCreateContext();
extern void WindstormMakeContextCurrent(WindstormWindow window, WindstormContext context);
extern void WindstormSwapBuffers(WindstormWindow window);
extern void WindstormStop();

extern char *errorMsg;

*/
import "C"
import "errors"

type cWindow C.WindstormWindow
type cContext C.WindstormContext

func cInit() error {

	_, err := C.WindstormInit()
	if err != nil {
		return errors.New(C.GoString(C.errorMsg))
	}

	return nil
}

//export resizeEvent
func resizeEvent(width, height int, window C.WindstormWindow) {

	if obj, ok := windows[cWindow(window)]; ok {
		// This function is called with every event check on X11, so manually
		// check to see if a change has indeed occurred.
		if obj.width != width || obj.height != height {
			obj.onResize(width, height)
		}
	}
}

//export keyboardEvent
func keyboardEvent(key int, action int, window C.WindstormWindow) {

	if obj, ok := windows[cWindow(window)]; ok {
		obj.onKeyboard(key, Action(action))
	}
}

//export mouseMoveEvent
func mouseMoveEvent(x, y int, window C.WindstormWindow) {

	if obj, ok := windows[cWindow(window)]; ok {
		obj.onMouseMove(x, y)
	}
}

func cNewWindow(width, height int, title string) (cWindow, error) {

	window, err := C.WindstormNewWindow(C.int(width), C.int(height), C.CString(title))
	if err != nil {
		return cWindow(window), errors.New(C.GoString(C.errorMsg))
	}

	return cWindow(window), nil
}

func cShowWindow(window cWindow) error {

	_, err := C.WindstormShowWindow(C.WindstormWindow(window))

	if err != nil {
		return errors.New(C.GoString(C.errorMsg))
	}

	return nil
}

func cHideWindow(window cWindow) error {

	_, err := C.WindstormHideWindow(C.WindstormWindow(window))

	if err != nil {
		return errors.New(C.GoString(C.errorMsg))
	}

	return nil
}

func cUpdateEvents(window cWindow) error {

	_, err := C.WindstormUpdateEvents(C.WindstormWindow(window))

	if err != nil {
		return errors.New(C.GoString(C.errorMsg))
	}

	return nil
}

func cCloseWindow(window cWindow) error {

	_, err := C.WindstormCloseWindow(C.WindstormWindow(window))

	if err != nil {
		return errors.New(C.GoString(C.errorMsg))
	}

	return nil
}

func cCreateContext() (cContext, error) {

	context, err := C.WindstormCreateContext()

	if err != nil {
		return cContext(context), errors.New(C.GoString(C.errorMsg))
	}

	return cContext(context), nil
}

func cMakeContextCurrent(window cWindow, context cContext) error {

	_, err := C.WindstormMakeContextCurrent(C.WindstormWindow(window), C.WindstormContext(context))

	if err != nil {
		return errors.New(C.GoString(C.errorMsg))
	}

	return nil
}

func cSwapBuffers(window cWindow) error {

	_, err := C.WindstormSwapBuffers(C.WindstormWindow(window))

	if err != nil {
		return errors.New(C.GoString(C.errorMsg))
	}

	return nil
}

func cStop() {

	C.WindstormStop()
}
