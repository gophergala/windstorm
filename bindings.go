package windstorm

/*
#include <X11/Xlib.h>
typedef Window WindstormWindow;

extern void WindstormInit();
extern WindstormWindow WindstormNewWindow(int, int, char*);
extern void WindstormShowWindow(WindstormWindow);
extern void WindstormHideWindow(WindstormWindow);
extern void WindstormUpdateEvents(WindstormWindow);

*/
import "C"
import "errors"

type cWindow C.WindstormWindow

func cInit() error {

	_, err := C.WindstormInit()
	if err != nil {
		return errors.New("could not initialize window system")
	}

	return nil
}

func cNewWindow(width, height int, title string) (cWindow, error) {

	window, err := C.WindstormNewWindow(C.int(width), C.int(height), C.CString(title))
	if err != nil {
		return cWindow(window), err
	}

	return cWindow(window), nil
}

func cShowWindow(window cWindow) error {

	_, err := C.WindstormShowWindow(C.WindstormWindow(window))

	if err != nil {
		return errors.New("failed to show window")
	}

	return nil
}

func cHideWindow(window cWindow) error {

	_, err := C.WindstormHideWindow(C.WindstormWindow(window))

	if err != nil {
		return errors.New("failed to hide window")
	}

	return nil
}

func cUpdateEvents(window cWindow) error {

	_, err := C.WindstormUpdateEvents(C.WindstormWindow(window))

	if err != nil {
		return errors.New("could not update events")
	}

	return nil
}
