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

## Document

[GoDoc](https://godoc.org/github.com/joway/trie)

## Benchmark

```
BenchmarkBuild-8                                    2000            814677 ns/op          555392 B/op       7377 allocs/op
BenchmarkTrie_AddWord-8                         100000000               12.6 ns/op             0 B/op          0 allocs/op
BenchmarkTrie_PrefixSearchString-8               5000000               267 ns/op              32 B/op          1 allocs/op
BenchmarkTrie_PrefixSearch-8                    10000000               156 ns/op               0 B/op          0 allocs/op
```

### Benchmark Comparison

[https://github.com/derekparker/trie](https://github.com/derekparker/trie) 

```
BenchmarkTrie_AddWord_DTrie-8                   20000000               108 ns/op             112 B/op          2 allocs/op
BenchmarkTrie_PrefixSearchString_DTrie-8         2000000               732 ns/op             160 B/op          9 allocs/op
```
