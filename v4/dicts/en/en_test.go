package en

import (
	"testing"

	"github.com/aaaton/golem/v4"
)

func TestEnglishUsage(t *testing.T) {
	l, err := golem.New(New())
	if err != nil {
		t.Fatal(err)
	}
	word := l.Lemma("agreed")
	// fmt.Println(word)
	result := "agree"
	if word != result {
		t.Errorf("Wanted %s '%b', got %s '%b'.", result, []byte(result), word, []byte(word))
	}

	word = l.Lemma("first")
	result = "1"
	if word != result {
		t.Errorf("Wanted %s '%b', got %s '%b'.", result, []byte(result), word, []byte(word))
	}
}
