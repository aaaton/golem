package main

import (
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

type localStorage struct {
	Lookup map[string]int
	Words  [][]string
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

	ls := &localStorage{}
	m := make(map[string][]string)
	for _, line := range strings.Split(strings.TrimSpace(string(b)), "\n") {
		parts := strings.Split(strings.TrimSpace(line), "\t")
		if len(parts) == 2 {
			base := strings.ToLower(parts[0])
			form := strings.ToLower(parts[1])
			add(m, form, base)
			add(m, base, base)
		} else {
			fmt.Printf("the line >%s< is odd\n", line)
		}
	}
	joined2Index := make(map[string]int)
	var forms []string
	ls.Lookup = make(map[string]int)
	for k, v := range m {
		lookup := strings.Join(v, "|")
		index, ok := joined2Index[lookup]
		if !ok {
			index = len(ls.Words)
			joined2Index[lookup] = index
			ls.Words = append(ls.Words, v)
		}
		forms = append(forms, k)
		ls.Lookup[k] = index
	}

	count := 0.0
	for _, form := range forms {
		if _, found := ls.Lookup[form]; !found {
			count++
		}
	}
	if count > 0 {
		fmt.Printf("Couldn't find %f%% of the keys entered\n", count/float64(len(forms))*100)
		os.Exit(1)
	}

	f, err = os.Create(outName)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = gob.NewEncoder(f).Encode(ls)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Words in dict:", len(forms))
	fmt.Println("Saved to", outName, "and all is good")
}

func add(m map[string][]string, key, value string) {
	if values, ok := m[key]; ok {
		if !contains(values, value) {
			values = append(values, value)
			sort.Strings(values)
			m[key] = values
		}
	} else {
		m[key] = []string{value}
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
