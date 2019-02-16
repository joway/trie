package trie

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func createTrie() *Trie {
	var dict = map[string]interface{}{
		"/abc": "2",
		"/a":   "1",
		"/ac":  "3",
		"/b":   "4",
		"/bc":  "5",
		"/bca": "6",
		"/ba":  "7",
		"/cba": "8",
	}
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
