package es

import (
	"fmt"
	"testing"

	"github.com/aaaton/golem/v4"
)

func TestSpanishUsage(t *testing.T) {
	l, err := golem.New(New())
	if err != nil {
		fmt.Println(err)
	}
	_ = l
	word := l.Lemma("primer")
	result := "1"
	if word != result {
		t.Errorf("Wanted %s '%b', got %s '%b'.", result, []byte(result), word, []byte(word))
	}

	word = l.Lemma("Buenas")
	result = "bueno"
	if word != result {
		t.Errorf("Wanted %s '%b', got %s '%b'.", result, []byte(result), word, []byte(word))
	}
}
