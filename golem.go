package golem

import (
	"fmt"
	"sort"
	"strings"
)

// LanguagePack is what each language should implement
type LanguagePack interface {
	GetResource() ([]byte, error)
	GetLocale() string
}

// Lemmatizer is the key to lemmatizing a word in a language
type Lemmatizer struct {
	m map[string]int
	v [][]string
}

type storage struct {
	Lookup map[string]int
	Words  [][]string
}

func newStorage(b []byte) (storage, error) {
	lines := strings.Split(string(b), "\n")
	s := storage{
		Lookup: make(map[string]int),
		Words:  [][]string{},
	}
	// TODO: Would it be better to do with a reader
	// instead of loading the full thing into an array?

	// br := bufio.NewReader(bytes.NewReader(b))
	// line, err := br.ReadString('\n')
	// for err == nil {
	// wordIndex := make(map[string])
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		words := strings.Split(line, "\t")
		if len(words) < 2 {
			return s, fmt.Errorf("expected more than 1 form per word")
		}
		base := words[0]
		for _, word := range words {
			if index, ok := s.Lookup[word]; ok {
				s.Words[index] = append(s.Words[index], word)
			} else {
				index := len(s.Words)
				s.Words = append(s.Words, []string{base})
				s.Lookup[word] = index
			}
		}
	}
	return s, nil
}

// New produces a new Lemmatizer
func New(pack LanguagePack) (*Lemmatizer, error) {
	resource, err := pack.GetResource()
	if err != nil {
		return nil, fmt.Errorf(`Could not open resource file for "%s"`, pack.GetLocale())
	}
	s, err := newStorage(resource)
	if err != nil {
		return nil, fmt.Errorf(`language %s is not valid: %s`, pack.GetLocale(), err)
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
// expects `word` to be lowercased
func (l *Lemmatizer) LemmaLower(word string) string {
	if out, ok := l.m[word]; ok {
		return l.v[out][0]
	}
	return word
}

// Lemmas gets all the base forms of a word, if multiple exist
func (l *Lemmatizer) Lemmas(word string) (out []string) {
	if index, ok := l.m[strings.ToLower(word)]; ok {
		out := l.v[index]
		// to get rid of the randomness, we sort the output
		sort.Strings(out)
		return out
	}
	return []string{word}
}
