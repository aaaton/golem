package golem

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"testing"

	"github.com/charlesgiroux/golem/dicts"
)

func TestReadBinary(t *testing.T) {
	b, err := dicts.Asset("data/en.gz")
	if err != nil {
		panic(err)
	}
	_ = b
	gzip.NewReader(bytes.NewBuffer(b))
}

func TestUsage(t *testing.T) {
	l, err := New("english")
	if err != nil {
		fmt.Println(err)
	}
	_ = l
	word := l.Lemma("agreed")
	fmt.Println(word)
	result := "agree"
	if word != result {
		t.Errorf("Wanted %s, got %s.", result, word)
	}
}

func TestFrenchUsage(t *testing.T) {
	l, err := New("fr")
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
	l, err := New("es")
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
	l, err := New("de")
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

func TestPrint(t *testing.T) {
	l, err := New("sv")
	if err != nil {
		panic(err)
	}
	added := make(map[string]bool)
	for k, v := range l.m {
		fmt.Println(k)
		added[k] = true
		for _, w := range v {
			if !added[w] {
				fmt.Println(w)
			}
			added[w] = true
		}
	}
}

func TestLemmatizer_Lemma(t *testing.T) {
	l, _ := New("swedish")
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
		})
	}
}

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New("swedish")
	}
}

func BenchmarkLookup(b *testing.B) {
	l, _ := New("swedish")
	for i := 0; i < b.N; i++ {
		l.Lemma("Avtalet")
	}
}
