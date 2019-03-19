package main

import (
	"fmt"
	"strings"
)

func main() {
	input := ` int(11) unsigned
 int(11) unsigned
 int(11) unsigned
 int(11) unsigned
 int(11) unsigned
 int(11) unsigned`
	split := strings.Split(input, "\n")
	r := make([]string, len(split))
	for i, s := range split {
		if strings.Contains(s, "bigint") {
			r[i] = "int64"
		} else if strings.Contains(s, "datetime") {
			r[i] = "time.Time"
		} else if strings.Contains(s, "varchar") {
			r[i] = "string"
		} else if strings.Contains(s, "tinyint") || strings.Contains(s, "int") {
			r[i] = "int"
		} else if strings.Contains(s, "double") {
			r[i] = "float64"
		}
		fmt.Println(r[i])
	}
}
