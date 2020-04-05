package golem

import (
	"strings"
	"testing"
)

//TestDict is a pretend LanguagePack for testing purposes
type TestDict struct {
}

//GetResource shows the expected dictionary format.
// 1 line per base form, with additional forms to the right. Tab separated
func (t *TestDict) GetResource() ([]byte, error) {
	return []byte(`word	wordy	wordis	wordlike
thing	thingy	thingies	thingimagics`), nil
}

func (t *TestDict) GetLocale() string {
	return "test"
}

func TestTestDict(t *testing.T) {
	dict := &TestDict{}
	lem, err := New(dict)
	if err != nil {
		t.Fatal(err)
	}
	b, _ := dict.GetResource()
	for _, line := range strings.Split(string(b), "\n") {
		words := strings.Split(line, "\t")
		base := words[0]
		for _, word := range words {
			lemma := lem.Lemma(word)
			if lemma != base {
				t.Errorf("Expected %v, but got %v for %v", base, lemma, word)
			}
		}
	}
}

func BenchmarkLookup(b *testing.B) {
	l, err := New(&TestDict{})
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N/2; i++ {
		l.Lemma("Thingimagics")
	}
}

func BenchmarkLookupLower(b *testing.B) {
	l, err := New(&TestDict{})
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N/2; i++ {
		l.LemmaLower("thingimagics")
	}
}
