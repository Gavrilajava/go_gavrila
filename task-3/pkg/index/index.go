package index

import (
	"go-gavrila/task-3/pkg/crawler"
	"strings"
	"sync"
)

var index = make(map[string][]int)
var content = []crawler.Document{}
var counter int
var mu sync.Mutex

// Adds documents to a collection and maintains indices.
func Add(links []crawler.Document) {
	for _, link := range links {
		mu.Lock()
		counter++
		link.ID = counter
		content = append(content, link)
		words := strings.Split(strings.ToLower(link.Title), " ")
		for _, word := range words {
			if not_indexed(word, link.ID) {
				index[word] = append(index[word], link.ID)
			}
		}
		mu.Unlock()
	}
}

// Retrieves documents based on a specified word. If the word is empty,
// it returns all the documents in the content collection.
// Result is sorted by url
func Collect(word string) []crawler.Document {
	if word == "" {
		return insertionSort(content)
	}
	var result []crawler.Document
	for _, id := range index[word] {
		result = append(result, find(id))
	}
	return insertionSort(result)
}

// Retrieves a document based on its ID.
func find(id int) crawler.Document {
	low, high := 0, len(content)-1
	for low <= high {
		mid := (low + high) / 2
		if content[mid].ID == id {
			return content[mid]
		}
		if content[mid].ID < id {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return crawler.Document{}
}

// Insertion sort algorithm implementation for crawler.Document
func insertionSort(arr []crawler.Document) []crawler.Document {

	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && strings.Compare(arr[j].URL, key.URL) > 0 {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}

	return arr
}

// Checks if url has already been indexed for the word.
func not_indexed(word string, id int) bool {
	if len(index[word]) > 0 {
		return index[word][len(index[word]) - 1] != id
	} else {
		return true
	}
}