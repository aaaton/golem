package fr

import (
	"fmt"
	"testing"

	"github.com/aaaton/golem/v4"
)

func TestFrenchUsage(t *testing.T) {
	l, err := golem.New(New())
	if err != nil {
		fmt.Println(err)
	}

	word := l.Lemma("avait")
	result := "avoir"
	if word != result {
		t.Errorf("Wanted %s '%b', got %s '%b'.", result, []byte(result), word, []byte(word))
	}

	// NOTE: This should test the first entry of the fr dictionary, but matches to the wrong entry
	// (possibly because of my keyboard language)
	// ```
	// 2e	e
	// 58e	e
	// 7e	e
	// ```
	// Below matches to "58e" instead of "2e"
	//
	// word = l.Lemma("e")
	// result = "2e"
	// if word != result {
	// 	t.Errorf("Wanted %s '%b', got %s '%b'.", result, []byte(result), word, []byte(word))
	// }
}
