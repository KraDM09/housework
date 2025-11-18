package memcached

import (
	"encoding/json"
	"github.com/bradfitz/gomemcache/memcache"
)

func (m *memcached) Set(key string, value any, expire int32) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return m.client.Set(&memcache.Item{
		Key:        KeyPrefix + key,
		Value:      data,
		Expiration: expire,
	})
}
