package main

import (
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	// ahocorasick "github.com/aaaton/Ahocorasick"
	slim "github.com/openacid/slim/trie"
)

type localStorage struct {
	Trie        *slim.Trie
	Words       [][]string
	indexLookup map[string]int
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage: cedar_creator [input] [output]")
		os.Exit(1)
	}
	inName, outName := os.Args[1], os.Args[2]
	f, err := os.Open(inName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ls := &localStorage{Cedar: ahocorasick.NewCedar(), Words: [][]string{}, indexLookup: make(map[string]int)}
	// cedar := ahocorasick.NewCedar()
	for i, line := range strings.Split(strings.TrimSpace(string(b)), "\n") {
		parts := strings.Split(strings.TrimSpace(line), "\t")
		if len(parts) == 2 {
			base := strings.ToLower(parts[0])
			form := strings.ToLower(parts[1])
			fmt.Println("inserting", i, form, base)
			insert(ls, []byte(form), base)
			fmt.Println("base")
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
	defer f.Close()
	// err = ls.Cedar.Save(f, "json")
	// if err != nil {
	// 	panic(err)
	// }
	dataEecoder := gob.NewEncoder(f)
	dataEecoder.Encode(ls)
	fmt.Println("Saved to", outName, "and all is good")
	// fmt.Println("Testing:")
	// words := []string{"spelled", "spld", "splendid", "use", "used"}
	// fmt.Println("\nBefore saving:")
	// for _, w := range words {
	// 	v, err := ls.Cedar.Get([]byte(w))
	// 	words := ls.Words[v.(int)]
	// 	if err != nil {
	// 		fmt.Println("Couldn't find", w)
	// 	} else {
	// 		fmt.Println("Found", words, "for", w)
	// 	}
	// }
	// var s storage
	// of, err := os.Open(outName)
	// if err != nil {
	// 	panic(err)
	// }
	// defer of.Close()
	// err = gob.NewDecoder(of).Decode(&s)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("\nAfter loading:")
	// for _, w := range words {
	// 	v, err := s.Cedar.Get([]byte(w))
	// 	words := s.Words[v.(int)]
	// 	if err != nil {
	// 		fmt.Println("Couldn't find", w)
	// 	} else {
	// 		fmt.Println("Found", words, "for", w)
	// 	}
	// }
}

type storage struct {
	Cedar *ahocorasick.Cedar
	Words [][]string
}

func insert(ls *localStorage, key []byte, value string) {
	v, err := ls.Cedar.Get(key)
	var words []string
	if err != nil {
		words = []string{value}
	} else { // key exists
		values := ls.Words[v.(int)]
		if !contains(values, value) {
			words = append(values, value)
			// insert = append(values, value)
		} else {
			words = values
		}
	}

	lookup := strings.Join(words, "|")
	if _, ok := ls.indexLookup[lookup]; !ok {
		ls.indexLookup[lookup] = len(ls.Words)
		ls.Words = append(ls.Words, words)
	}
	fmt.Println("cedar insert")
	err = ls.Cedar.Insert(key, ls.indexLookup[lookup])
	if err != nil {
		panic(err)
	}
}

func contains(values []string, value string) bool {
	for _, v := range values {
		if v == value {
			return true
		}
	}
	return false
}
