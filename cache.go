package cache

import "errors"

type CacheItem any

type Cache struct {
	data map[string]CacheItem
}

func NewCache() *Cache {
	return &Cache{data: make(map[string]CacheItem)}
}

func (c *Cache) Get(key string) (any, error) {
	value, exists := c.data[key]

	if !exists {
		return nil, errors.New("Value does not exist")
	}

    return value, nil
}

func (c *Cache) Set(key string, value any) {
    c.data[key] = value
}

func (c *Cache) Delete(key string) (any, error) {
    value, exists := c.data[key]

    if !exists {
        return nil, errors.New("Value not found")
    }

    delete(c.data, key)

    return value, nil
}
