package main

import (
	"fmt"
	"io/ioutil"
)


func main() {
	txt, err := ioutil.ReadFile("test.txt")
	check (err)
	fmt.Println(string(txt))
}


func check(err error){
	if err != nil {
		fmt.Println(err)
	}
}