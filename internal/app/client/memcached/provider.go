package memcached

import "github.com/bradfitz/gomemcache/memcache"

const KeyPrefix = "housework_"

type memcached struct {
	client *memcache.Client
}

func NewMemcached(server string) (Provider, error) {
	mc := memcache.New(server)

	err := mc.Ping()
	if err != nil {
		return nil, err
	}

	return &memcached{
		client: mc,
	}, nil
}

type Provider interface {
	Set(key string, value any, expire int32) error
	Get(key string) ([]byte, error)
}
