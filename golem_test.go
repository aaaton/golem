package golem

import (
	"fmt"
	"strings"
	"testing"

	"github.com/aaaton/golem/dicts/de"
	"github.com/aaaton/golem/dicts/en"
	"github.com/aaaton/golem/dicts/es"
	"github.com/aaaton/golem/dicts/fr"
	"github.com/aaaton/golem/dicts/sv"
)

func TestEnglishUsage(t *testing.T) {
	l, err := New(en.New())
	if err != nil {
		t.Fatal(err)
	}
	word := l.Lemma("agreed")
	fmt.Println(word)
	result := "agree"
	if word != result {
		t.Errorf("Wanted %s, got %s.", result, word)
	}
}

func TestFrenchUsage(t *testing.T) {
	l, err := New(fr.New())
	if err != nil {
		fmt.Println(err)
	}

	word := l.Lemma("avait")
	fmt.Println(word)
	result := "avoir"
	if word != result {
		t.Errorf("Wanted %s, got %s.", result, word)
	}
}

func TestSpanishUsage(t *testing.T) {
	l, err := New(es.New())
	if err != nil {
		fmt.Println(err)
	}
	_ = l
	word := l.Lemma("Buenas")
	fmt.Println(word)
	result := "bueno"
	if word != result {
		t.Errorf("Wanted %s, got %s.", result, word)
	}
}

func TestGermanUsage(t *testing.T) {
	l, err := New(de.New())
	if err != nil {
		fmt.Println(err)
	}
	_ = l
	word := l.Lemma("Hast")
	fmt.Println(word)
	result := "haben"
	if word != result {
		t.Errorf("Wanted %s, got %s.", result, word)
	}
}

func TestLemmatizer_Lemma(t *testing.T) {
	l, err := New(sv.New())
	if err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		in  string
		out string
	}{
		{"Avtalet", "avtal"},
		{"avtalets", "avtal"},
		{"avtalens", "avtal"},
		{"Avtaletsadlkj", "Avtaletsadlkj"},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			got := l.Lemma(tt.in)
			if got != tt.out {
				t.Errorf("Lemmatizer.Lemma() = %v, want %v", got, tt.out)
			}
			got = l.LemmaLower(strings.ToLower(tt.in))
			if got != strings.ToLower(tt.out) {
				t.Errorf("Lemmatizer.LemmaLower() = %v, want %v", got, tt.out)
			}
		})
	}
}

func BenchmarkLookup(b *testing.B) {
	l, err := New(sv.New())
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N/2; i++ {
		l.Lemma("Avtalet")
	}
}

func BenchmarkLookupLower(b *testing.B) {
	l, err := New(sv.New())
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N/2; i++ {
		l.LemmaLower("avtalet")
	}
}
