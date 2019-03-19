package main

import "fmt"

func main() {
	result := make([]int, 0)
	fmt.Println(result)
	for i := 1; i <= 3; i++ {
		result = append(result, i)
	}
	fmt.Println(result)
}
