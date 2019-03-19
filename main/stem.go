package main

import (
	"fmt"
	"github.com/aaaton/golem"
	"io/ioutil"
	"os"
	"strings"
)

type Word struct {
	word          string
	pronunciation string
	chinese       string
}

var lemmatizer *golem.Lemmatizer

func init() {
	lemmatizerL, err := golem.New("english")
	if err != nil {
		panic(err)
	}
	lemmatizer = lemmatizerL
}
func main() {
	//
	//word := lemmatizer.Lemma("conclusions")
	//fmt.Println(word)

	bytes, e := ioutil.ReadFile("main/dict")
	if e != nil {
		panic(e)
	}

	dict := make(map[string]*Word)

	for _, s := range strings.Split(string(bytes), "\n") {
		split := strings.Split(s, "\t")
		w := &Word{}
		w.word = split[0]
		if len(split) == 2 {
			w.chinese = split[1]
		} else {
			w.pronunciation = split[1]
			w.chinese = strings.Join(split[2:], " ")
		}
		dict[w.word] = w
	}

	file, e := ioutil.ReadFile("main/input")
	if e != nil {
		panic(e)
	}
	content := string(file)

	result := ""
	punctuation := []rune{',', '.', '?', '!'}

	for _, s := range strings.Split(content, " ") {
		st := s
		st = strings.TrimSpace(st)
		st = strings.TrimFunc(st, func(r rune) bool {
			for _, p := range punctuation {
				if r == p {
					return true
				}
			}
			return false
		})
		w, ok := dict[lemmatizer.Lemma(st)]
		if ok {
			result += s + fmt.Sprintf("[%s] ", w.chinese)
		} else {
			result += s + " "
		}
	}
	e = ioutil.WriteFile("main/output", []byte(result), os.ModePerm)
	if e != nil {
		panic(e)
	}

	fmt.Println("done")

	//fmt.Println(dict)
}
