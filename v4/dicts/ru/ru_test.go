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
}
