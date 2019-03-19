package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	end := time.Date(2019, 03, 17, 8, 0, 0, 0, time.Local)
	sub := start.Sub(end)
	fmt.Println(sub)
	fmt.Println(int(sub.Minutes()))

}
