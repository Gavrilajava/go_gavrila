package index

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"strings"

	"go-gavrila/task-13/pkg/crawler"
	"go-gavrila/task-13/pkg/crawler/spider"
)

const storage_path = `./storage.json`
const depth = 2

var urls = [2]string{"https://go.dev", "https://golang.org"}

type Service struct {
	Index     map[string][]int   `json:"index"`
	Documents []crawler.Document `json:"Documents"`
	Counter   int                `json:"counter"`
}

func New() (*Service, error) {

	s := Service{
		Index:     make(map[string][]int),
		Documents: []crawler.Document{},
	}

	f, err := os.Open(storage_path)
	if err != nil {
		if os.IsNotExist(err) {
			// if file does not exist, just return empty service
			return &s, nil
		} else {
			return nil, err
		}
	}
	defer f.Close()

	s, err = read(f)
	if err != nil {
		return nil, err
	}

	if s.Empty() {
		spider := spider.New()
		for _, url := range urls {
			res, err := spider.Scan(url, depth)
			if err != nil {
				continue
			}
			for _, link := range res {
				s.Add(link)
			}
		}
		if err = s.Save(); err != nil {
			return nil, err
		}
	}

	return &s, nil

}

func (s Service) Empty() bool {
	return len(s.Documents) < 1
}

// Adds documents to the collection and maintains indices.
func (s *Service) Add(link crawler.Document) crawler.Document {
	s.Counter++
	link.ID = s.Counter
	s.Documents = append(s.Documents, link)
	words := strings.Split(strings.ToLower(link.Title), ` `)
	for _, word := range words {
		if s.not_indexed(word, link.ID) {
			s.Index[word] = append(s.Index[word], link.ID)
		}
	}
	return link
}

// Updates document in the collection
func (s *Service) Update(d crawler.Document, params crawler.Document) crawler.Document {
	params.ID = d.ID
	s.Documents[s.FindIndex(d.ID)] = params
	words := strings.Split(strings.ToLower(d.Title), ` `)
	for _, word := range words {
		if s.not_indexed(word, d.ID) {
			s.Index[word] = append(s.Index[word], d.ID)
		}
	}
	return params
}

// Deletes document from the collection
func (s *Service) Delete(d crawler.Document) []crawler.Document {
	i := s.FindIndex(d.ID)
	s.Documents = append(s.Documents[:i], s.Documents[i+1:]...)
	return s.Documents
}

func (s Service) Save() error {
	f, err := os.Create(storage_path)
	if err != nil {
		return err
	}
	defer f.Close()
	return s.write(f)
}

// Retrieves documents based on a specified word. If the word is empty,
// it returns all the documents in the Documents collection.
// Result is sorted by url
func (s Service) Collect(word string) []crawler.Document {
	if word == "" {
		return insertionSort(s.Documents)
	}
	var result []crawler.Document
	for _, id := range s.Index[word] {
		result = append(result, s.Documents[s.FindIndex(id)])
	}
	return insertionSort(result)
}

func (s Service) Find(id int) (crawler.Document, error) {
	i := s.FindIndex(id)
	if i < 0 {
		return crawler.Document{}, errors.New("document not found")
	}
	return s.Documents[i], nil
}

// Retrieves index of a document based on its ID.
func (s Service) FindIndex(id int) int {
	low, high := 0, len(s.Documents)-1
	for low <= high {
		mid := (low + high) / 2
		if s.Documents[mid].ID == id {
			return mid
		}
		if s.Documents[mid].ID < id {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

// Checks if url has already been indexed for the word.
func (s Service) not_indexed(word string, id int) bool {
	if len(s.Index[word]) > 0 {
		return s.Index[word][len(s.Index[word])-1] != id
	} else {
		return true
	}
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

// reads data from provided source
func read(r io.Reader) (Service, error) {

	var s Service

	data, err := io.ReadAll(r)
	if err != nil {
		return s, err
	}

	err = json.Unmarshal(data, &s)
	if err != nil {
		return s, err
	}
	return s, nil
}

// writes a sting to the writer
func (s Service) write(w io.Writer) error {

	j, err := json.Marshal(s)
	if err != nil {
		return err
	}

	_, err = w.Write(j)
	if err != nil {
		return err
	}
	return nil
}
