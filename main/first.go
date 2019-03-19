package main

import "fmt"

func main() {

	var a *bool
	a = new(bool(true))
	fmt.Println(a)
}
