package trie

import (
	"fmt"
	"sort"
	"strings"
)

var rootKey rune = -1

type Trie struct {
	Key   rune
	Value interface{}
	//use map to improve search performance
	Next map[rune]*Trie
}

func New(key rune, value interface{}) *Trie {
	return &Trie{
		Key:   key,
		Value: value,
		Next:  map[rune]*Trie{},
	}
}

func Build(dict map[string]interface{}) *Trie {
	root := New(rootKey, nil)
	for w, v := range dict {
		root.AddWord([]rune(w), v)
	}

	return root
}

func (t *Trie) IsEnd() bool {
	return len(t.Next) == 0
}

func (t *Trie) IsRoot() bool {
	return t.Key == rootKey
}

func (t *Trie) AddWord(word []rune, value interface{}) {
	if len(word) == 0 {
		return
	}
	c := word[0]
	next := t.Next[c]
	if next == nil {
		next = New(c, nil)
		t.Next[c] = next
	}
	if len(word) == 1 {
		next.Value = value
	}

	next.AddWord(word[1:], value)
}

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
	} else {
		return maxDepth + 1
	}
}

func (t *Trie) PrefixSearchString(word string) (string, interface{}) {
	w, v := t.PrefixSearch([]rune(word))
	return string(w), v
}

func (t *Trie) PrefixSearch(word []rune) ([]rune, interface{}) {
	if len(word) == 0 {
		return nil, nil
	}

	hits := make([]rune, 0)
	var val interface{} = nil
	next := t
	for _, w := range word {
		next = next.Next[w]
		if next == nil {
			break
		}
		hits = append(hits, w)
		val = next.Value
	}

	return hits, val
}

func (t *Trie) Prettify() string {
	words := extractWords(t)
	output := ""
	lastWord := ""
	for index, word := range words[1:] {
		if index > 0 {
			lastWord = words[index-1]
		}
		curLine := ""
		for i, c := range word {
			curLine += string(c)
			content := ""
			isFirst := i == 0
			if strings.HasPrefix(lastWord, curLine) {
				spaceCount := 2
				if isFirst {
					spaceCount = 0
				}
				content += "|" + strings.Repeat(" ", spaceCount)
			} else {
				if !isFirst {
					content += "->"
				}

				content += string(c)
				lastWord += string(c)
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
