package golem

import (
	"bytes"
	"compress/gzip"
	"strings"
	"testing"

	"github.com/aaaton/golem/dicts"
)

func TestReadBinary(t *testing.T) {
	b, err := dicts.Asset("data/en.gz")
	if err != nil {
		t.Fatal(err)
	}
  _, err = gzip.NewReader(bytes.NewBuffer(b))
  if err != nil {
		t.Fatal(err)
	}
}

func TestUsage(t *testing.T) {
	l, err := New("english")
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

func TestLemmatizer_Lemma(t *testing.T) {
	l, err := New("swedish")
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
