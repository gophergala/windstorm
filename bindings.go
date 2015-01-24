package windstorm

/*
#include <X11/Xlib.h>
typedef Window WindstormWindow;

extern void WindstormInit();
extern WindstormWindow WindstormNewWindow(int, int, char*);
extern void WindstormShowWindow(WindstormWindow);
extern void WindstormHideWindow(WindstormWindow);
extern void WindstormUpdateEvents(WindstormWindow);
extern void WindstormCloseWindow(WindstormWindow);

extern char *errorMsg;

*/
import "C"
import "errors"

type cWindow C.WindstormWindow

func cInit() error {

	_, err := C.WindstormInit()
	if err != nil {
		return errors.New(C.GoString(C.errorMsg))
	}

	return nil
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
