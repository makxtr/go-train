package main

import "fmt"

type WordCounter struct {
	order  []string
	counts map[string]int
	limit  int
}

func NewWordCounter(limit int) *WordCounter {
	return &WordCounter{
		order:  make([]string, 0),
		counts: make(map[string]int),
		limit:  limit,
	}
}

func (wc *WordCounter) CountWord(word string) {
	if _, ok := wc.counts[word]; !ok {
		wc.order = append(wc.order, word)
	}
	wc.counts[word]++

	if len(wc.counts) > wc.limit {
		delete(wc.counts, wc.order[0])
		wc.order = wc.order[1:]
	}
}

func main() {
	wc := NewWordCounter(3)

	words := []string{"apple", "banana", "apple", "orange", "grape", "banana", "kiwi"}
	for _, word := range words {
		wc.CountWord(word)
	}

	fmt.Println("Количество слов:", wc.counts, wc.order)
}
