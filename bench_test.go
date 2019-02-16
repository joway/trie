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
	dict := createBenchDict()
	b.ResetTimer()
	t := Build(dict)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		word := "福建龙海市石码"
		t.PrefixSearchString(word)
	}
}

func BenchmarkTrie_PrefixSearch(b *testing.B) {
	dict := createBenchDict()
	b.ResetTimer()
	t := Build(dict)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		word := "福建龙海市石码"
		t.PrefixSearch([]rune(word))
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
