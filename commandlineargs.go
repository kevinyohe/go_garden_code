package main

import (
	"fmt"
	"os"
)


func main(){
	for i,v:= range os.Args{
		fmt.Printf("Arguments recieved: %v %v\n", i, v)
	}
}