package trie

import (
	"io/ioutil"
	"math/rand"
	"strings"
	"testing"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func createBenchDict() map[string]interface{} {
	dict := map[string]interface{}{}
	content, _ := ioutil.ReadFile("./words.txt")
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		k := line
		v := randString(1)
		dict[k] = v
	}
	return dict
}

func BenchmarkTrie_PrefixSearchString(b *testing.B) {
	b.ReportAllocs()
	dict := createBenchDict()
	b.ResetTimer()
	t := Build(dict)
	for i := 0; i < b.N; i++ {
		word := "福建龙海市石码"
		t.PrefixSearchString(word)
	}
}

func BenchmarkTrie_PrefixSearch(b *testing.B) {
	b.ReportAllocs()
	dict := createBenchDict()
	b.ResetTimer()
	t := Build(dict)
	for i := 0; i < b.N; i++ {
		word := "福建龙海市石码"
		t.PrefixSearch([]rune(word))
	}
}

func BenchmarkTrie_AddWord(b *testing.B) {
	b.ReportAllocs()
	dict := createBenchDict()
	root := New(rootKey, nil)
	count := len(dict)
	words := make([][]rune, count)
	for w := range dict {
		words = append(words, []rune(w))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		root.AddWord(words[i%count], "xxx")
	}
}

func BenchmarkBuild(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dict := createBenchDict()
		Build(dict)
	}
	dict := createBenchDict()
	Build(dict)
}
