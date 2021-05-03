package main

import "fmt"


func add4(p *int) {
	*p = 4 + *p
}
func add4noptr(p int) int{
	return p+4
}
// * value of
// & address of
func main(){
	x := 40
	add4(&x)
	fmt.Println(x)
	fmt.Println(add4noptr(x))
}
