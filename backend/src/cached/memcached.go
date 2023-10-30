package cached

import (
	"time"

	"github.com/bradfitz/gomemcache/memcache"
)

type Memcached memcache.Client

func NewMemcached(endpoint ...string) *Memcached {
	return (*Memcached)(memcache.New(endpoint...))
}

func (m *Memcached) Save(key string, value []byte) error {
	return m.SaveFor(0, key, value)
}

func (m *Memcached) SaveFor(expire time.Duration, key string, value []byte) error {
	c := (*memcache.Client)(m)
	return c.Add(&memcache.Item{
		Key:        key,
		Value:      value,
		Expiration: int32(expire.Seconds()),
	})
}

func (m *Memcached) Get(key string) (value []byte, err error) {
	c := (*memcache.Client)(m)
	if item, err := c.Get(key); err != nil {
		if err != memcache.ErrCacheMiss {
			return nil, nil
		} else {
			return nil, err
		}
	} else {
		return item.Value, nil
	}
}