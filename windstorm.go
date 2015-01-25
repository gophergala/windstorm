// Package windstorm is a window creation library for Go.
package windstorm

func init() {

	windows = make(map[cWindow]*Window)

	err := cInit()
	if err != nil {
		panic(err)
	}
}

// Stop stops the library. This should only be called if no more Windstorm
// tasks will be done for the duration of the program.
func Stop() {

	cStop()
}
