package golem

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"sort"
	"strings"

	"github.com/aaaton/golem/dicts"
)

// Lemmatizer is the key to lemmatizing a word in a language
type Lemmatizer struct {
	m map[string]int
	v [][]string
}

type storage struct {
	Lookup map[string]int
	Words  [][]string
}

// New produces a new Lemmatizer
func New(pack dicts.LanguagePack) (*Lemmatizer, error) {
	resource, err := pack.GetResource()
	if err != nil {
		return nil, fmt.Errorf(`Could not open resource file for "%s"`, pack.GetLocale())
	}
	var s storage
	err = gob.NewDecoder(bytes.NewBuffer(resource)).Decode(&s)
	if err != nil {
		return nil, fmt.Errorf(`language %s is not valid`, pack.GetLocale())
	}
	l := Lemmatizer{m: s.Lookup, v: s.Words}
	return &l, nil
}

// InDict checks if a certain word is in the dictionary
func (l *Lemmatizer) InDict(word string) bool {
	_, ok := l.m[strings.ToLower(word)]
	return ok
}

// Lemma gets one of the base forms of a word
func (l *Lemmatizer) Lemma(word string) string {
	if out, ok := l.m[strings.ToLower(word)]; ok {
		return l.v[out][0]
	}
	return word
}

// LemmaLower gets one of the base forms of a lower case word
func (l *Lemmatizer) LemmaLower(word string) string {
	if out, ok := l.m[word]; ok {
		return l.v[out][0]
	}
	return word
}

// Lemmas gets all the base forms of a word
func (l *Lemmatizer) Lemmas(word string) (out []string) {
	if index, ok := l.m[strings.ToLower(word)]; ok {
		out := l.v[index]
		// to get rid of the randomness, we sort the output
		sort.Strings(out)
		return out
	}
	return []string{word}
}
