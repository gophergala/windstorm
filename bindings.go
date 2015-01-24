package windstorm

/*
extern void WindstormInit();
*/
import "C"
import "errors"

func WingolInit() error {

	_, err := C.WindstormInit()
	if err != nil {
		return errors.New("could not initialize window system")
	}

	return nil
}
