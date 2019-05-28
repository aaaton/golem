package main

import (
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	ahocorasick "github.com/aaaton/Ahocorasick"
)

type localStorage struct {
	Cedar       *ahocorasick.Cedar
	Words       []string
	indexLookup map[string]int
}

func main() {
	if len(os.Args) != 3 {
		panic("usage: cedar_creator [input] [output]")
	}
	inName, outName := os.Args[1], os.Args[2]
	f, err := os.Open(inName)
	if err != nil {
		panic(err)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	ls := &localStorage{Cedar: ahocorasick.NewCedar(), Words: []string{}, indexLookup: make(map[string]int)}
	// cedar := ahocorasick.NewCedar()
	for _, line := range strings.Split(string(b), "\n") {
		parts := strings.Split(strings.TrimSpace(line), "\t")
		if len(parts) == 2 {
			base := strings.ToLower(parts[0])
			form := []byte(strings.ToLower(parts[1]))
			insert(ls, form, base)
			insert(ls, []byte(base), base)
		} else {
			fmt.Printf("the line >%s< is odd\n", line)
		}
	}
	keys, nodes, size, capacity := ls.Cedar.Status()
	fmt.Printf("Keys: %v, nodes:%v, size: %v, capacity: %v\n", keys, nodes, size, capacity)
	fmt.Printf("Words: %v\n", len(ls.Words))
	f, err = os.Create(outName)
	if err != nil {
		panic(err)
	}
	// err = ls.Cedar.Save(f, "json")
	// if err != nil {
	// 	panic(err)
	// }
	dataEecoder := gob.NewEncoder(f)
	dataEecoder.Encode(ls)
	fmt.Println("Saved to", outName, "and all is good")
	// words := []string{"spelled", "spld", "splendid", "use", "used"}
	// for _, w := range words {
	// 	v, err := cedar.Get([]byte(w))
	// 	if err != nil {
	// 		fmt.Println("Couldn't find", w)
	// 	} else {
	// 		fmt.Println("Found", v, "for", w)
	// 	}
	// }
	// cedar.Insert(key []byte, value interface{})
}

func insert(ls *localStorage, key []byte, value string) {
	index, ok := ls.indexLookup[value]
	if !ok {
		ls.indexLookup[value] = len(ls.Words)
		ls.Words = append(ls.Words, value)
	}
	v, err := ls.Cedar.Get(key)
	var indexes []int
	if err != nil {
		indexes = []int{index}
	} else { // key exists
		indexes = v.([]int)
		if !contains(indexes, index) {
			indexes = append(indexes, index)
		}
	}
	err = ls.Cedar.Insert(key, indexes)
	if err != nil {
		panic(err)
	}
}

func contains(indexes []int, index int) bool {
	for _, v := range indexes {
		if v == index {
			return true
		}
	}
	return false
}
