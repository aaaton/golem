package main_test

import (
	"math/rand"
	// "sort"
	"testing"

	"github.com/derekparker/trie"
	dghubble "github.com/dghubble/trie"
)

func TestDerekParkerTrie(t *testing.T) {
	var keys []string
	tr := trie.New()
	for i := 0; i < 100000; i++ {
		key := randStringRunes(10)
		keys = append(keys, key)
		tr.Add(key, i)
	}
	count := 0
	failed := 0
	for i, key := range keys {
		n, found := tr.Find(key)
		if !found {
			count++
		} else if n.Meta().(int) != i {
			failed++
		}
	}
	if count > 0 {
		t.Errorf("Couldn't find %f%% of the keys entered. Missing %d keys", float64(count)/float64(len(keys))*100, count)
	}
	if failed > 0 {
		t.Errorf("Got the wrong result for %d keys", failed)
	}
}

func BenchmarkDerekParkerGet(b *testing.B) {
	keys := getKeys(100000)
	tr := trie.New()
	for i := 0; i < len(keys); i++ {
		tr.Add(keys[i], i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tr.Find(keys[i%len(keys)])
	}
}

func TestDGHubbleTrie(t *testing.T) {
	tr, keys := createHubble(100000)

	count := 0
	failed := 0
	for i, key := range keys {
		out := tr.Get(key)
		if out == nil {
			count++
		} else if out.(int) != i {
			failed++
		}
	}
	if count > 0 {
		t.Errorf("Couldn't find %f%% of the keys entered. Missing %d keys", float64(count)/float64(len(keys))*100, count)
	}
	if failed > 0 {
		t.Errorf("Got the wrong result for %d keys", failed)
	}
}
func createHubble(num int) (*dghubble.PathTrie, []string) {
	keys := getKeys(num)
	tr := dghubble.NewPathTrie()
	for i := 0; i < len(keys); i++ {
		tr.Put(keys[i], i)
	}
	return tr, keys
}
func BenchmarkHubbleGet(b *testing.B) {
	tr, keys := createHubble(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tr.Get(keys[i%len(keys)])
	}
}

func getKeys(num int) (keys []string) {
	for i := 0; i < num; i++ {
		key := randStringRunes(10)
		keys = append(keys, key)
		// indexes = append(indexes, i)
	}
	return
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
