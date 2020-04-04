package main

import (
	"bufio"
	"compress/zlib"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

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

	m := make(map[string][]string)
	for _, line := range strings.Split(strings.TrimSpace(string(b)), "\n") {
		parts := strings.Split(strings.TrimSpace(line), "\t")
		if len(parts) == 2 {
			base := strings.ToLower(parts[0])
			form := strings.ToLower(parts[1])
			add(m, base, form)
			// add(m, base, base)
		} else {
			fmt.Printf("the line >%s< is odd\n", line)
		}
	}

	f, err = os.Create(outName)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	zw, err := zlib.NewWriterLevel(f, zlib.BestCompression)
	if err != nil {
		panic(err)
	}
	defer zw.Close()
	wr := bufio.NewWriter(zw)
	for k, values := range m {
		_, err = wr.WriteString(strings.Join(append([]string{k}, values...), "\t"))
		if err != nil {
			panic(err)
		}
		_, err = wr.WriteRune('\n')
		if err != nil {
			panic(err)
		}
	}
	wr.Flush()
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
