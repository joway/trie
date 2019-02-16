# Trie

[![Go Report Card](https://goreportcard.com/badge/github.com/joway/trie)](https://goreportcard.com/report/github.com/joway/trie)
[![codecov](https://codecov.io/gh/joway/trie/branch/master/graph/badge.svg)](https://codecov.io/gh/joway/trie)
[![CircleCI](https://circleci.com/gh/joway/trie.svg?style=shield)](https://circleci.com/gh/joway/trie)

## Usage

```go
import "github.com/joway/trie"

dict := map[string]interface{}{
		"/abc": "2",
		"/a":   "1",
		"/ac":  "3",
		"/b":   "4",
		"/bc":  "5",
		"/bca": "6",
		"/ba":  "7",
		"/cba": "8",
}

tree := trie.Build(dict)
prefix, val := tree.PrefixSearchString("/a")
```
## Benchmark

```
BenchmarkTrie_PrefixSearchString-8       5000000               362 ns/op              88 B/op          4 allocs/op
BenchmarkTrie_PrefixSearch-8             5000000               255 ns/op              56 B/op          3 allocs/op
BenchmarkBuild-8                            2000            827605 ns/op
```
