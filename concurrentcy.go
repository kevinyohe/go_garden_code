package main

import (
	"fmt"
	"time"
)

func count(item string){
	for i:=1; i<=3; i++{
		fmt.Printf("%v %v  ", i, item)
		time.Sleep(1 * time.Second)
	}
	fmt.Println()
}


func main(){
	fmt.Println("\nLine-by-line Execution:")
	count("Moose")
	count("sheep")

	fmt.Println("\nConcurrent Execution:")
	go count("Moose")
	count("sheep")
}