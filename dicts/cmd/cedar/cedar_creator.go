package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"

	// ahocorasick "github.com/aaaton/Ahocorasick"
	"github.com/openacid/slim/encode"
	slim "github.com/openacid/slim/trie"
)

type localStorage struct {
	Trie  *slim.SlimTrie
	Words [][]string
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage: cedar_creator [input] [output]")
		os.Exit(1)
	}
	inName, _ := os.Args[1], os.Args[2]
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

	ls := &localStorage{Words: [][]string{}}
	// slim.NewSlimTrie(encode.Int, keys []string, values interface{})
	// cedar := ahocorasick.NewCedar()
	m := make(map[string][]string)
	for _, line := range strings.Split(strings.TrimSpace(string(b)), "\n") {
		parts := strings.Split(strings.TrimSpace(line), "\t")
		if len(parts) == 2 {
			base := strings.ToLower(parts[0])
			form := strings.ToLower(parts[1])
			if values, ok := m[form]; ok {
				if !contains(values, base) {
					m[form] = append(values, base)
				}
			} else {
				if form == "used" {
					fmt.Println("Adding", form, base)
				}
				m[form] = []string{base}
			}
			if values, ok := m[base]; ok {
				if !contains(values, base) {
					m[base] = append(values, base)
				}
			} else {
				m[base] = []string{base}
			}
		} else {
			fmt.Printf("the line >%s< is odd\n", line)
		}
	}
	fmt.Println("Map contains used?", m["used"])
	joined2Index := make(map[string]int)
	var forms []string
	form2index := make(map[string]int)
	for k, v := range m {
		lookup := strings.Join(v, "|")
		index, ok := joined2Index[lookup]
		if !ok {
			index = len(ls.Words)
			joined2Index[lookup] = index
			ls.Words = append(ls.Words, v)
		}
		forms = append(forms, k)
		form2index[k] = index
		// indexes = append(indexes, index)
	}
	fmt.Println("form2index contains used:", form2index["used"], ls.Words[form2index["used"]])
	sort.Strings(forms)
	var indexes []int
	for _, form := range forms {
		if form == "used" {
			fmt.Println("Forms contains used:", form, form2index[form])
		}
		indexes = append(indexes, form2index[form])
	}
	ls.Trie, err = slim.NewSlimTrie(encode.Int{}, forms, indexes)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Words: %v\n", len(ls.Words))
	// f, err = os.Create(outName)
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()
	// // err = ls.Cedar.Save(f, "json")
	// // if err != nil {
	// // 	panic(err)
	// // }
	// dataEecoder := gob.NewEncoder(f)
	// dataEecoder.Encode(ls)
	// fmt.Println("Saved to", outName, "and all is good")
	count := 0.0
	for _, form := range forms {
		if _, found := ls.Trie.Get(form); !found {
			// fmt.Println("Couldn't find", form)
			count++
		}
	}
	fmt.Println(len(forms))
	fmt.Printf("Couldn't find %f%% of the keys entered\n", count/float64(len(forms))*100)
	fmt.Println("\nTesting:")
	words := []string{"spelled", "spld", "splendid", "use", "used"}
	fmt.Println("Before saving:")
	for _, w := range words {
		v, found := ls.Trie.Get(w)
		if !found {
			fmt.Println("Couldn't find", w)
		} else {
			words := ls.Words[v.(int)]
			fmt.Println("Found", words, "for", w)
		}
	}
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

// func insert(ls *localStorage, key []byte, value string) {
// 	v, err := ls.Cedar.Get(key)
// 	var words []string
// 	if err != nil {
// 		words = []string{value}
// 	} else { // key exists
// 		values := ls.Words[v.(int)]
// 		if !contains(values, value) {
// 			words = append(values, value)
// 			// insert = append(values, value)
// 		} else {
// 			words = values
// 		}
// 	}

// 	lookup := strings.Join(words, "|")
// 	if _, ok := ls.indexLookup[lookup]; !ok {
// 		ls.indexLookup[lookup] = len(ls.Words)
// 		ls.Words = append(ls.Words, words)
// 	}
// 	fmt.Println("cedar insert")
// 	err = ls.Cedar.Insert(key, ls.indexLookup[lookup])
// 	if err != nil {
// 		panic(err)
// 	}
// }

func contains(values []string, value string) bool {
	for _, v := range values {
		if v == value {
			return true
		}
	}
	return false
}
