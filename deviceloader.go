package main

import (
	"awesomeProject/devicedata"
	"fmt"
	"time"
)





func main() {
	fmt.Println("Hello, playground")
	a := devicedata.DeviceData{Name: "test123", Command: "show ip br", Data: ""}
	fmt.Println(a.Name)
	a.Name = "new name"
	fmt.Println(a.Name)
	a.Save()
	fmt.Print(time.Now())
	fmt.Println()
	dt:= time.Now()
	_, _, d := dt.Date()
	fmt.Println(d, dt.Weekday())
}

