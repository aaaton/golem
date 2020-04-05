package de

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
	word := l.Lemma("Hast")
	result := "haben"
	if word != result {
		t.Errorf("Wanted %s, got %s.", result, word)
	}
}
