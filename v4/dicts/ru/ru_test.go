package ru

import (
	"testing"

	"github.com/aaaton/golem/v4"
)

func TestRussianUsage(t *testing.T) {
	l, err := golem.New(New())
	if err != nil {
		t.Fatal(err)
	}
	word := l.Lemma("масла")
	// fmt.Println(word)
	result := "масло"
	if word != result {
		t.Errorf("Wanted %s, got %s.", result, word)
	}

	word = l.Lemma("человека")
	result = "человек"
	if word != result {
		t.Errorf("Wanted %s '%b', got %s '%b'.", result, []byte(result), word, []byte(word))
	}
}
