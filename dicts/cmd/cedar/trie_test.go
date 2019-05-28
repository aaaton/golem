package main_test

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/openacid/slim/encode"
	"github.com/openacid/slim/trie"
)

func TestSlimTrie(t *testing.T) {
	var keys []string
	var indexes []int
	for i := 0; i < 70000; i++ {
		keys = append(keys, randStringRunes(10))
		indexes = append(indexes, i)
	}
	sort.Strings(keys)
	tr, err := trie.NewSlimTrie(encode.Int{}, keys, indexes)
	if err != nil {
		t.Error(err)
		return
	}
	count := 0
	for _, key := range keys {
		_, found := tr.Get(key)
		if !found {
			count++
		}
	}
	if count > 0 {
		t.Errorf("Couldn't find %f%% of the keys entered. Missing %d keys", float64(count)/float64(len(keys))*100,count)
	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
