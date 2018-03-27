package logstash

import (
	"errors"
	"fmt"
	"sync"

	"github.com/go-redis/redis"
)

type redisSyncer struct {
	once   sync.Once
	client *redis.Client
	config *Config
}

func (s *redisSyncer) createClient() error {
	var retErr error
	s.once.Do(func() {
		c := redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", s.config.Host, s.config.Port),
			Password: s.config.Password,
			DB:       s.config.DB,
			PoolSize: 20,
		})

		err := c.Ping().Err()
		if err != nil {
			retErr = err
		} else {
			s.client = c
		}
	})

	return retErr
}

func (s *redisSyncer) Write(p []byte) (n int, err error) {
	if err := s.createClient(); err != nil {
		return 0, err
	}

	if s.config.DataType == List {
		v, e := s.client.RPush(s.config.Key, p).Result()
		n = int(v)
		return n, e
	}

	if s.config.DataType == Channel {
		v, e := s.client.Publish(s.config.Key, p).Result()
		n = int(v)
		return n, e
	}

	return 0, errors.New("Unknow data type")
}
