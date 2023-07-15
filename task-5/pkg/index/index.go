package index

import (
	"bufio"
	"fmt"
	"go-gavrila/task-5/pkg/crawler"
	"io"
	"os"
	"strconv"
	"strings"
)

const content_path = `./content.csv`
const index_path = `./index.csv`
const separator = `;`

var index = make(map[string][]int)
var content = []crawler.Document{}

func LoadFiles() {
	load_content(load_file(content_path))
	load_index(load_file(index_path))
}

// Loads data from file and converts it to strings slice
func load_file(name string) []string {
	fmt.Println("loading", name)
	file, err := os.Open(name)
	if err != nil {
		fmt.Println(err)
		return []string{}
	}
	defer file.Close()

	return read(file)

}

// reads data from provided source
func read(r io.Reader) []string {

	reader := bufio.NewReader(r)
	var lines []string

	for line, err := reader.ReadString('\n'); err != nil {
		// line, err := reader.ReadString('\n')
		// if err != nil {
		// 	break
		// }
		lines = append(lines, strings.TrimSuffix(line, "\n"))
	}

	fmt.Println("imported", len(lines), "items")

	return lines
}

// writes a sting to the writer
func write(w io.Writer, s string) error {
	_, err := w.Write([]byte(s))
	return err
}

// Converts slice of strings to an array of crawler.Document
func load_content(data []string) {

	for _, line := range data {

		arr := strings.Split(line, separator)
		i, err := strconv.ParseInt(arr[0], 0, 64)
		if err != nil {
			fmt.Println(err)
		} else {
			content = append(content, crawler.Document{int(i), arr[1], arr[2], arr[3]})
		}

	}

}

// Converts strings slice to a map
func load_index(data []string) {

	for _, line := range data {
		arr := strings.Split(line, separator)
		for _, i := range strings.Split(arr[1], `,`) {
			v, err := strconv.ParseInt(i, 0, 64)
			if err != nil {
				fmt.Println(err)
			} else {
				index[arr[0]] = append(index[arr[0]], int(v))
			}
		}
	}

}

// Adds documents to a collection and maintains indices.
func Add(links []crawler.Document) {
	var counter int
	if len(content) > 0 {
		counter = content[len(content)-1].ID
		// Clear the old index from file if we have previous content
		if err := os.Truncate(index_path, 0); err != nil {
			fmt.Println(err)
			return
		}
	}

	content_storage, err := os.OpenFile(content_path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(content_storage.Name())

	for _, link := range links {
		counter++
		link.ID = counter
		content = append(content, link)
		if err := write(content_storage, fmt.Sprintf("%d;%s;%s;%s\n", link.ID, link.URL, link.Title, link.Body)); err != nil {
			fmt.Println(err)
		}
		words := strings.Split(strings.ToLower(link.Title), ` `)
		for _, word := range words {
			if not_indexed(word, link.ID) {
				index[word] = append(index[word], link.ID)
			}
		}
	}
	defer content_storage.Close()

	index_storage, err := os.OpenFile(index_path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	for word, indices := range index {
		if err := write(index_storage, fmt.Sprintf("%s;%s\n", word, strings.Trim(strings.Replace(fmt.Sprint(indices), " ", `,`, -1), "[]"))); err != nil {
			fmt.Println(err)
		}
	}
	defer index_storage.Close()
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
		return index[word][len(index[word])-1] != id
	} else {
		return true
	}
}

func Empty() bool {
	return len(content) < 1
}
