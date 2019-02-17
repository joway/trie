package trie

import (
	"fmt"
	"sort"
	"strings"
)

var rootKey rune = -1

// Trie is the trie node struct
type Trie struct {
	Key   rune
	Value interface{}
	//use map to improve search performance
	Next map[rune]*Trie
}

// New create a Trie Node
func New(key rune, value interface{}) *Trie {
	return &Trie{
		Key:   key,
		Value: value,
		Next:  map[rune]*Trie{},
	}
}

func NewRoot() *Trie {
	return New(rootKey, nil)
}

// Build build a new Trie from a map
func Build(dict map[string]interface{}) *Trie {
	root := NewRoot()
	for w, v := range dict {
		root.AddWordString(w, v)
	}

	return root
}

// IsEnd check is the end of the trie
func (t *Trie) IsEnd() bool {
	return len(t.Next) == 0
}

// IsRoot check is the root of the trie
func (t *Trie) IsRoot() bool {
	return t.Key == rootKey
}

// AddWord add a word into the trie
func (t *Trie) AddWord(word []rune, value interface{}) {
	parent := t
	for _, c := range word {
		if parent.Next[c] == nil {
			parent.Next[c] = New(c, nil)
		}
		parent = parent.Next[c]
	}
	parent.Value = value
}

// AddWordString add a string word into the trie
func (t *Trie) AddWordString(word string, value interface{}) {
	t.AddWord([]rune(word), value)
}

// Depth get the depth of the trie
func (t *Trie) Depth() int {
	maxDepth := 0
	for _, child := range t.Next {
		d := child.Depth()
		if maxDepth < d {
			maxDepth = d
		}
	}
	if t.IsRoot() {
		return maxDepth
	}
	return maxDepth + 1
}

// Forward get the match trie node from word
func (t *Trie) Forward(word []rune) *Trie {
	next := t
	for _, w := range word {
		if next.Next[w] == nil {
			break
		}
		next = next.Next[w]
	}
	return next
}

// PrefixSearchString prefix search by a string word
func (t *Trie) PrefixSearchString(word string) (string, interface{}) {
	w, v := t.PrefixSearch([]rune(word))
	return string(w), v
}

// PrefixSearch prefix search by a []rune word
func (t *Trie) PrefixSearch(word []rune) ([]rune, interface{}) {
	prefixIndex := -1
	var val interface{}
	next := t
	for i, w := range word {
		next = next.Next[w]
		if next == nil {
			break
		}
		prefixIndex = i
		val = next.Value
	}

	if prefixIndex < 0 {
		return []rune{}, nil
	}

	return word[0 : prefixIndex+1], val
}

// Prettify get the prettify graph of the trie
func (t *Trie) Prettify() string {
	words := extractWords(t)
	output := ""
	lastWord := ""
	for index, word := range words {
		if index > 0 {
			lastWord = words[index-1]
		}
		curWord := ""
		for i, c := range word {
			char := string(c)
			curWord += char
			content := ""
			isFirst := i == 0
			if strings.HasPrefix(lastWord, curWord) {
				spaceCount := 2
				if isFirst {
					spaceCount = 0
				}
				content += "|" + strings.Repeat(" ", spaceCount)
			} else {
				if !isFirst {
					content += "->"
				}

				content += char
				lastWord += char
			}

			output += content
		}
		output += "\n"
	}
	return output
}

func extractWords(t *Trie) []string {
	//DFS
	if t == nil {
		return []string{}
	}

	curKey := t.Key
	if t.IsRoot() {
		curKey = []rune("*")[0]
	}
	words := make([]string, 0)
	nextKeys := keys(t.Next)
	//let all keys ordered
	sort.SliceStable(nextKeys, func(i, j int) bool {
		return nextKeys[i] < nextKeys[j]
	})

	for _, key := range nextKeys {
		subTrie := t.Next[key]
		suffixs := extractWords(subTrie)
		for _, suffix := range suffixs {
			words = append(words, fmt.Sprintf("%s%s", string(curKey), suffix))
		}
	}
	if len(words) == 0 {
		return []string{string(curKey)}
	}
	return words
}

func keys(dict map[rune]*Trie) []rune {
	ks := make([]rune, len(dict))
	for k := range dict {
		ks = append(ks, k)
	}
	return ks
}
