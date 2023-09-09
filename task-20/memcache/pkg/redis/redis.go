package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type Record struct {
	Short       string    `json:"short"`
	Long  			string    `json:"long"`
}

type Storage struct {
	client *redis.Client
}

func New(addr string, password string, db int) (*Storage, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return &Storage{
		client: client,
	}, nil
}

func (storage *Storage) Add(r Record) error {
	err := storage.client.Set(context.Background(), r.Short, r.Long, 24 * time.Hour).Err()
	if err != nil {
		return err
	}

	return nil
}

func (storage *Storage) Get(short string) (Record, error) {
	long, err := storage.client.Get(context.Background(), short).Result()
	if err != nil {
		return Record{}, err
	}

	return Record{
		Short: short,
		Long:  long,
	}, nil
}