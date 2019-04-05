package golem

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"fmt"
	"path"
	"strings"

	"github.com/aaaton/golem/dicts"
	it "github.com/axamon/golem/dicts/IT"
)

// Lemmatizer is the key to lemmatizing a word in a language
type Lemmatizer struct {
	m map[string][]string
}

const folder = "data"

// NewVoid produces a void Lemmatizer
func NewVoid() (*Lemmatizer, error) {
	l := Lemmatizer{m: make(map[string][]string)}
	return &l, nil

}

// New produces a new Lemmatizer
func New(locale string) (*Lemmatizer, error) {
	var fname string

	switch locale {
	case "sv", "swedish":
		fname = "sv.gz"
	case "en", "english":
		fname = "en.gz"
	case "fr", "french":
		fname = "fr.gz"
	case "de", "german":
		fname = "de.gz"
	case "es", "spanish":
		fname = "es.gz"
	case "it", "italian":
		fname = "it.gz"
	default:
		return nil, fmt.Errorf(`Language "%s" is not implemented`, locale)
	}

	var f []byte
	var err error

	switch fname {
	case "it.gz":
		f, err = it.Asset(path.Join(folder, fname))
	default:
		f, err = dicts.Asset(path.Join(folder, fname))
	}

	if err != nil {
		return nil, fmt.Errorf(`Could not open resource file for "%s"`, locale)
	}
	r, err := gzip.NewReader(bytes.NewBuffer(f))
	if err != nil {
		return nil, fmt.Errorf(`Could not open resource file for "%s"`, locale)
	}

	l := Lemmatizer{m: make(map[string][]string)}
	br := bufio.NewReader(r)
	line, err := br.ReadString('\n')
	for err == nil {
		parts := strings.Split(strings.TrimSpace(line), "\t")
		if len(parts) == 2 {
			base := strings.ToLower(parts[0])
			form := strings.ToLower(parts[1])
			l.m[form] = append(l.m[form], base)
			l.m[base] = append(l.m[base], base)
		} else {
			fmt.Println(line, "is odd")
		}
		line, err = br.ReadString('\n')
	}
	return &l, nil
}

// InDict checks if a certain word is in the dictionary
func (l *Lemmatizer) InDict(word string) bool {
	_, ok := l.m[strings.ToLower(word)]
	return ok
}

// Lemma gets one of the base forms of a word
func (l *Lemmatizer) Lemma(word string) string {
	loweredword := strings.ToLower(word)
	if out, ok := l.m[loweredword]; ok {
		return out[0]
	}
	return word
}

// LemmaLower gets one of the base forms of a lower case word
func (l *Lemmatizer) LemmaLower(word string) string {
	if out, ok := l.m[word]; ok {
		return out[0]
	}
	return word
}

// Lemmas gets all the base forms of a word
func (l *Lemmatizer) Lemmas(word string) []string {
	if out, ok := l.m[strings.ToLower(word)]; ok {
		return out
	}
	return []string{word}
}
