package it

import (
	"testing"

	"github.com/aaaton/golem/v4"
)

func TestItalian(t *testing.T) {
	lem, err := golem.New(New())
	if err != nil {
		t.Fatal(err)
	}
	// abbarbicare	abbarbicai
	word := lem.Lemma("abbarbicai")
	result := "abbarbicare"
	if word != result {
		t.Errorf("Wanted %s '%b', got %s '%b'.", result, []byte(result), word, []byte(word))
	}

	word = lem.Lemma("abati")
	result = "abate"
	if word != result {
		t.Errorf("Wanted %s '%b', got %s '%b'.", result, []byte(result), word, []byte(word))
	}
}
