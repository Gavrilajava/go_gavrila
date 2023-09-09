package cache

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"context"

	"shortener/pkg/storage"
)

type Service struct {
	path string
}

func New(path string) *Service {
	return &Service{
		path: path,
	}
}


func (c *Service) Add(ctx context.Context, r storage.Record) error {
	body, err := json.Marshal(r)

	if err != nil {
		return err
	}

	client := http.Client{}
	req , err := http.NewRequest("POST", c.path, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("requestId", fmt.Sprintf("%s", ctx.Value("requestId")))

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

func (c *Service) Get(ctx context.Context, short string) (*storage.Record, error) {

	client := http.Client{}
	req , err := http.NewRequest("GET", c.path + "/" + short, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("requestId", fmt.Sprintf("%s", ctx.Value("requestId")))

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	record := storage.Record{}

	err = json.NewDecoder(resp.Body).Decode(&record)
	if err != nil {
		return nil, err
	}

	return &record, nil
}