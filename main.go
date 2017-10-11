package golem

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"fmt"
	"path"
	"strings"
)

// Lemmatizer is the key to lemmatizing a word in a language
type Lemmatizer struct {
	m map[string]string
}

const folder = "data"

// New produces a new Lemmatizer
func New(locale string) (*Lemmatizer, error) {
	var fname string

	switch locale {
	case "sv", "swedish":
		fname = "sv.gz"
	case "en", "english":
		fname = "en.gz"
	default:
		return nil, fmt.Errorf(`Language "%s" is not implemented`, locale)
	}
	f, err := Asset(path.Join(folder, fname))
	if err != nil {
		return nil, fmt.Errorf(`Could not open resource file for "%s"`, locale)
	}
	r, err := gzip.NewReader(bytes.NewBuffer(f))
	if err != nil {
		return nil, fmt.Errorf(`Could not open resource file for "%s"`, locale)
	}

	l := Lemmatizer{m: make(map[string]string)}
	br := bufio.NewReader(r)
	line, err := br.ReadString('\n')
	for err == nil {
		parts := strings.Split(strings.TrimSpace(line), "\t")
		if len(parts) == 2 {
			l.m[strings.ToLower(parts[1])] = strings.ToLower(parts[0])
			l.m[strings.ToLower(parts[0])] = strings.ToLower(parts[0])
		} else {
			fmt.Println(line, "is odd")
		}
		line, err = br.ReadString('\n')
	}
	return &l, nil
}

// Lemma gets the base form of a word
func (l *Lemmatizer) Lemma(word string) (string, error) {
	if out, ok := l.m[strings.ToLower(word)]; ok {
		return out, nil
	}
	return "", fmt.Errorf("Word not found in dictionary")
}
