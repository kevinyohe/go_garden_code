package main

import (
	"fmt"
	"strings"
)

func main () {
	fmt.Println("test")
	strings.EqualFold("a", "A")
	str := "Everyone has a plan until they get punched in the mouth"
	fmt.Printf("String has prefix xxx: %v \n", strings.HasPrefix(str, "xxx"))
	fmt.Printf("String has suffix: %v \n", strings.HasSuffix(str,"outh"))
}