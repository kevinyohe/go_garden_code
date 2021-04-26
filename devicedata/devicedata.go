package devicedata

import "fmt"

type DeviceData struct {
	Name string
	Command string
	Data string
}

func (d DeviceData) Save() {
	fmt.Println("saving" , d.Name)
}
