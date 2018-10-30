package main

import "fmt"

func main() {
	a := int(1.5e6)
	fmt.Printf("%T, %v", a, a+1)
}
