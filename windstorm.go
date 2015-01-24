// Package windstorm is a window creation library for Go.
package windstorm

func init() {

	err := WingolInit()
	if err != nil {
		panic(err)
	}
}
