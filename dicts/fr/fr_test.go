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
		t.Errorf("Wanted %s, got %s.", result, word)
	}
}
