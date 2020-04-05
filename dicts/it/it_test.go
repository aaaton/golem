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
	result := "abbarbicare"
	lemma := lem.Lemma("abbarbicai")
	if lemma != result {
		t.Errorf("Expected %v, but got %v", result, lemma)
	}

}
