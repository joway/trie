package trie

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createDict() map[string]interface{} {
	return map[string]interface{}{
		"/abc": "2",
		"/a":   "1",
		"/ac":  "3",
		"/b":   "4",
		"/bc":  "5",
		"/bca": "6",
		"/ba":  "7",
		"/cba": "8",
	}
}

func createTrie() *Trie {
	var dict = createDict()
	t := Build(dict)
	return t
}

func TestTrie_PrefixSearch(t *testing.T) {
	tree := createTrie()
	path, val := tree.PrefixSearchString("/a")
	assert.Equal(t, "/a", path)
	assert.Equal(t, "1", val)

	path, val = tree.PrefixSearchString("/abc")
	assert.Equal(t, "/abc", path)
	assert.Equal(t, "2", val)

	path, val = tree.PrefixSearchString("/bccccc")
	assert.Equal(t, "/bc", path)
	assert.Equal(t, "5", val)
}

func TestTrie_Depth(t *testing.T) {
	tree := createTrie()
	depth := tree.Depth()
	assert.Equal(t, 4, depth)
}

func TestTrie_Forward(t *testing.T) {
	tree := createTrie()
	word := []rune("/abc")
	got := tree.Forward(word)
	assert.True(t, got.Key == word[len(word)-1])
	assert.True(t, got.IsEnd())
}

func TestTrie_Prettify(t *testing.T) {
	dict := createBenchDict()
	tree := Build(dict)
	output := tree.Prettify()
	for word := range dict {
		for _, c := range word {
			assert.Contains(t, output, string(c))
		}
	}
	fmt.Println(output)
}
