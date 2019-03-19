package main

import (
	"fmt"
	"strings"
)

func main() {

	source := toArray(`first_on_late_seconds
	first_off_early_seconds
	second_on_late_seconds
	second_off_early_seconds
	third_on_late_seconds
	third_off_early_seconds`)
	for _, v := range source {
		split := strings.Split(v, "_")
		r := make([]string, len(split))
		for i, w := range split {
			r[i] = strings.Title(w)
		}
		fmt.Println(strings.Join(r, ""))
	}

}

func toArray(s string) []string {
	split := strings.Split(s, "\n")
	r := make([]string, len(split))
	for i, w := range split {
		r[i] = strings.TrimSpace(w)
	}
	return r
}
