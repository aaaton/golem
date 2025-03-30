package uk

import (
	"testing"

	"github.com/aaaton/golem/v4"
)

func TestRussianUsage(t *testing.T) {
	l, err := golem.New(New())
	if err != nil {
		t.Fatal(err)
	}
	word := l.Lemma("автоматичная")
	// fmt.Println(word)
	result := "автоматичний"
	if word != result {
		t.Errorf("Wanted %s, got %s.", result, word)
	}

	word = l.Lemma("абзаца")
	result = "абзац"
	if word != result {
		t.Errorf("Wanted %s '%b', got %s '%b'.", result, []byte(result), word, []byte(word))
	}
}
