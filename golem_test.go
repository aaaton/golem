package golem_test

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/axamon/golem"
)

var en *golem.Lemmatizer = new(golem.Lemmatizer)
var fr *golem.Lemmatizer = new(golem.Lemmatizer)
var ge *golem.Lemmatizer = new(golem.Lemmatizer)
var it *golem.Lemmatizer = new(golem.Lemmatizer)
var sp *golem.Lemmatizer = new(golem.Lemmatizer)
var sw *golem.Lemmatizer = new(golem.Lemmatizer)

func TestMain(m *testing.M) {
	// Creates all lemmizators
	var err error
	en, err = golem.New("english")
	fr, err = golem.New("french")
	ge, err = golem.New("german")
	it, err = golem.New("italian")
	sp, err = golem.New("spanish")
	sw, err = golem.New("swedish")
	if err != nil {
		log.Fatalf("Impossibile creare Lemmatizer\n")
	}
	flag.Parse()

	exitCode := m.Run()

	// Exit
	os.Exit(exitCode)
}

var flagtests = []struct {
	name       string
	language   string
	Lemmatizer *golem.Lemmatizer
	in         string
	out        string
}{

	{"Swedish Example1", "swedish", sw, "Avtalet", "avtal"},
	{"Swedish Example2", "swedish", sw, "avtalets", "avtal"},
	{"Swedish Example3", "swedish", sw, "avtalens", "avtal"},
	{"Swedish Example4", "swedish", sw, "Avtaletsadlkj", "Avtaletsadlkj"},
	{"English Verb", "english", en, "goes", "go"},
	{"English Noun", "english", en, "wolves", "wolf"},
	{"English FirstName", "english", en, "Edward", "Edward"},
	{"French Example1", "french", fr, "avait", "avoir"},
	{"Spanish Example1", "spanish", sp, "Buenas", "bueno"},
	{"German Example1", "german", ge, "Hast", "haben"},
	{"Italian Verb", "italian", it, "lavorerai", "lavorare"},
	{"Italian Plural Noun", "italian", it, "bicchieri", "bicchiere"},
	{"Italian FirstName", "italian", it, "Alberto", "Alberto"},
	{"Italian Plural Adjective", "italian", it, "lunghi", "lungo"},
}

func TestLemmatizer_Lemma(t *testing.T) {
	for _, tt := range flagtests {
		t.Run(tt.name, func(t *testing.T) {
			var l *golem.Lemmatizer = new(golem.Lemmatizer)
			switch tt.language {
			case "italian":
				l = it
			case "english":
				l = en
			case "swedish":
				l = sw
			case "french":
				l = fr
			case "german":
				l = ge
			case "spanish":
				l = sp
			}
			var got string
			got = l.Lemma(tt.in)
			if got != tt.out {
				t.Errorf("%s Lemmatizer.Lemma() = %v, want %v", tt.name, got, tt.out)
			}
			got = l.LemmaLower(strings.ToLower(tt.in))
			if got != strings.ToLower(tt.out) {
				t.Errorf("%s Lemmatizer.LemmaLower() = %v, want %v", tt.name, got, tt.out)
			}
		})
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
		l, err := golem.New(element.language)
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
		var l *golem.Lemmatizer = new(golem.Lemmatizer)
		switch element.language {
		case "italian":
			l = it
		case "english":
			l = en
		case "swedish":
			l = sw
		case "french":
			l = fr
		case "german":
			l = ge
		case "spanish":
			l = sp
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
		var l *golem.Lemmatizer = new(golem.Lemmatizer)
		switch element.language {
		case "italian":
			l = it
		case "english":
			l = en
		case "swedish":
			l = sw
		case "french":
			l = fr
		case "german":
			l = ge
		case "spanish":
			l = sp
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
	l, err := golem.New("swedish")
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N/2; i++ {
		l.Lemma("Avtalet")
	}
}

func BenchmarkLookupLower(b *testing.B) {
	l, err := golem.New("swedish")
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N/2; i++ {
		l.LemmaLower("avtalet")
	}
}
