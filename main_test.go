package golem

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"log"
	"strings"
	"sync"
	"testing"

	it "github.com/axamon/golem/dicts/IT"
)

var flagtests = []struct {
	name     string
	language string
	in       string
	out      string
}{
	{"Italian Verb", "italian", "lavorerai", "lavorare"},
	{"Italian Plural Noun", "italian", "bicchieri", "bicchiere"},
	{"Italian FirstName", "italian", "Alberto", "Alberto"},
	{"Italian Plural Adjective", "italian", "lunghi", "lungo"},
	{"Swedish Example1", "swedish", "Avtalet", "avtal"},
	{"Swedish Example2", "swedish", "avtalets", "avtal"},
	{"Swedish Example3", "swedish", "avtalens", "avtal"},
	{"Swedish Example4", "swedish", "Avtaletsadlkj", "Avtaletsadlkj"},
	{"English Verb", "english", "goes", "go"},
	{"English Noun", "english", "wolves", "wolf"},
	{"English FirstName", "english", "Edward", "Edward"},
	{"French Example1", "french", "avait", "avoir"},
	{"Spanish Example1", "spanish", "Buenas", "bueno"},
	{"German Example1", "german", "Hast", "haben"},
}

func TestLemmatizer_Lemma_All(t *testing.T) {
	var wg sync.WaitGroup
	for _, tt := range flagtests {
		t.Run(tt.in, func(t *testing.T) {
			wg.Add(1)
			defer wg.Done()
			l, err := New(tt.language)
			if err != nil {
				t.Fatal(err)
			}
			//go func() {

			//lock.Lock()
			got := l.Lemma(tt.in)
			if got != tt.out {
				t.Errorf("%s Lemmatizer.Lemma() = %v, want %v", tt.name, got, tt.out)
			}
			//lock.Unlock()
			//lock.Lock()
			got = l.LemmaLower(strings.ToLower(tt.in))
			if got != strings.ToLower(tt.out) {
				t.Errorf("%s Lemmatizer.LemmaLower() = %v, want %v", tt.name, got, tt.out)
			}
			//lock.Unlock()
			//runtime.Gosched()
			//}()
		})
	}
	wg.Wait()
}

func TestReadBinary_IT(t *testing.T) {
	b, err := it.Asset("data/it.gz")
	if err != nil {
		t.Fatal(err)
	}
	_, err = gzip.NewReader(bytes.NewBuffer(b))
	if err != nil {
		t.Fatal(err)
	}
}

var exampleDataLemma = []struct {
	language string
	word     string
}{
	{"english", "agreed"},
	{"italian", "armadi"},
	{"swedish", "Avtalet"},
}

func ExampleLemmatizer_Lemma() {
	for _, element := range exampleDataLemma {
		l, err := New(element.language)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(l.Lemma(element.word))
	}
	// Output:
	// agree
	// armadio
	// avtal
}

var exampleDataInDict = []struct {
	language string
	word     string
	result   bool
}{
	{"italian", "armadio", true},
	{"italian", "ammaccabanane", false},
	{"swedish", "Avtalet", true},
	{"swedish", "Avtalt", false},
}

func ExampleLemmatizer_InDict() {
	for _, element := range exampleDataInDict {
		l, err := New(element.language)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(l.InDict(element.word))
	}
	// Output:
	// true
	// false
	// true
	// false
}

var exampleDataLemmas = []struct {
	language string
	word     string
	result   []string
}{
	{"italian", "soli", []string{"sole", "solo"}},
}

func ExampleLemmatizer_Lemmas() {
	for _, element := range exampleDataLemmas {
		l, err := New(element.language)
		if err != nil {
			log.Fatal(err)
		}
		lemmas := l.Lemmas(element.word)
		for _, lemma := range lemmas {
			fmt.Println(lemma)
		}
	}
	// Unordered output:
	// solare
	// solere
	// solo
	// sole
}

func BenchmarkLookup(b *testing.B) {
	l, err := New("swedish")
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N/2; i++ {
		l.Lemma("Avtalet")
	}
}

func BenchmarkLookupLower(b *testing.B) {
	l, err := New("swedish")
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N/2; i++ {
		l.LemmaLower("avtalet")
	}
}
