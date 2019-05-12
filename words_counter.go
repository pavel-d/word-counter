package main

import (
	"sort"
)

func NewWordsCounter(limit int) *WordsCounter {
	return &WordsCounter{tree: NewPrefixTree(), limit: limit}
}

type WordsCounter struct {
	tree  *PrefixTree
	limit int
}

func (w *WordsCounter) ProcessChunk(chunk []byte) {
	splitWords(chunk, func(word []byte) {
		w.tree.Insert(word)
	})
}

func (w *WordsCounter) Result() Dictionary {
	dict := w.tree.ListAll()
	sort.Sort(sort.Reverse(dict))
	return dict
}

func splitWords(text []byte, processFunc func([]byte)) {
	var wasLetter bool
	var wordStart int

	for idx, c := range text {
		if !isLetter(c) {
			if wasLetter {
				processFunc(toLower(text[wordStart:idx]))
				wasLetter = false
			}
		} else {
			if !wasLetter {
				wordStart = idx
				wasLetter = true
			}
		}
	}

	if wasLetter {
		processFunc(toLower(text[wordStart:len(text)]))
	}
}

func isLetter(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func isSpace(c byte) bool {
	return !isLetter(c)
}

func toLower(contents []byte) []byte {
	for idx, c := range contents {
		if !(c >= 'a' && c <= 'z') {
			contents[idx] += 'a' - 'A'
		}
	}

	return contents
}
