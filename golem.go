package golem

import (
	"bufio"
	"bytes"
	"fmt"
	"sort"
	"strings"
	// ahocorasick "github.com/BobuSumisu/go-ahocorasick"
)

// Lemmatizer is the key to lemmatizing a word in a language
type Lemmatizer struct {
	m map[string][]*string
}

// LanguagePack is what each language should implement
type LanguagePack interface {
	GetResource() ([]byte, error)
	GetLocale() string
}

// New produces a new Lemmatizer
func New(pack LanguagePack) (*Lemmatizer, error) {
	resource, err := pack.GetResource()
	if err != nil {
		return nil, fmt.Errorf(`Could not open resource file for "%s"`, pack.GetLocale())
	}
	l := Lemmatizer{m: make(map[string][]*string)}
	pointMap := make(map[string]*string)
	br := bufio.NewReader(bytes.NewBuffer(resource))
	line, err := br.ReadString('\n')
	for err == nil {
		parts := strings.Split(strings.TrimSpace(line), "\t")
		if len(parts) == 2 {
			base := strings.ToLower(parts[0])
			b, ok := pointMap[base]
			if !ok {
				b = &base
				pointMap[base] = b
			}
			form := strings.ToLower(parts[1])
			l.m[form] = append(l.m[form], b)
			l.m[base] = append(l.m[base], b)
		} else {
			println(line, "is odd")
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
	if out, ok := l.m[strings.ToLower(word)]; ok {
		return *out[0]
	}
	return word
}

// LemmaLower gets one of the base forms of a lower case word
func (l *Lemmatizer) LemmaLower(word string) string {
	if out, ok := l.m[word]; ok {
		return *out[0]
	}
	return word
}

// Lemmas gets all the base forms of a word
func (l *Lemmatizer) Lemmas(word string) (out []string) {
	if lemmas, ok := l.m[strings.ToLower(word)]; ok {
		for _, l := range lemmas {
			out = append(out, *l)
		}
		// to get rid of the randomness, we sort the output
		sort.Strings(out)
		return out
	}
	return []string{word}
}
