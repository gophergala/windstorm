// Package windstorm is a window creation library for Go.
package windstorm

func init() {

	windows = make(map[cWindow]*Window)

	err := cInit()
	if err != nil {
		panic(err)
	}
}

func Stop() {

	cStop()
}
