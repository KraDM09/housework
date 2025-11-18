package memcached

import (
	"errors"
	"github.com/bradfitz/gomemcache/memcache"
)

func (m *memcached) Get(key string) ([]byte, error) {
	row, err := m.client.Get(KeyPrefix + key)
	if errors.Is(err, memcache.ErrCacheMiss) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return row.Value, err
}
