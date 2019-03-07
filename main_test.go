package golem

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"testing"

	"github.com/axamon/golem/dicts"
)

func TestReadBinary(t *testing.T) {
	b, err := dicts.Asset("data/it.gz")
	if err != nil {
		t.Fatal(err)
	}
	_, err = gzip.NewReader(bytes.NewBuffer(b))
	if err != nil {
		t.Fatal(err)
	}
}

func TestUsage(t *testing.T) {
	l, err := New("italian")
	if err != nil {
		t.Fatal(err)
	}
	word := l.Lemma("patate")
	fmt.Println(word)
	result := "patata"
	if word != result {
		t.Errorf("volevo %s, e ho avuto %s.", result, word)
	}
}

func BenchmarkLookup(b *testing.B) {
	l, err := New("italian")
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N/2; i++ {
		l.Lemma("Attivo")
	}
}

func BenchmarkLookupLower(b *testing.B) {
	l, err := New("italian")
	if err != nil {
		b.Error(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N/2; i++ {
		l.LemmaLower("confetto")
	}
}
