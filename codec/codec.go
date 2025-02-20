package codec

import (
	"github.com/eko/gocache/store"
)

// Stats allows to returns some statistics of codec usage
type Stats struct {
	Hits              int
	Miss              int
	SetSuccess        int
	SetError          int
	DeleteSuccess     int
	DeleteError       int
	InvalidateSuccess int
	InvalidateError   int
}

// Codec represents an instance of a cache store
type Codec struct {
	store store.StoreInterface
	stats *Stats
}

// New return a new codec instance
func New(store store.StoreInterface) *Codec {
	return &Codec{
		store: store,
		stats: &Stats{},
	}
}

// Get allows to retrieve the value from a given key identifier
func (c *Codec) Get(key interface{}) (interface{}, error) {
	val, err := c.store.Get(key)

	if err == nil {
		c.stats.Hits++
	} else {
		c.stats.Miss++
	}

	return val, err
}

// Set allows to set a value for a given key identifier and also allows to specify
// an expiration time
func (c *Codec) Set(key interface{}, value interface{}, options *store.Options) error {
	err := c.store.Set(key, value, options)

	if err == nil {
		c.stats.SetSuccess++
	} else {
		c.stats.SetError++
	}

	return err
}

// Delete allows to remove a value for a given key identifier
func (c *Codec) Delete(key interface{}) error {
	err := c.store.Delete(key)

	if err == nil {
		c.stats.DeleteSuccess++
	} else {
		c.stats.DeleteError++
	}

	return err
}

// Invalidate invalidates some cach items from given options
func (c *Codec) Invalidate(options store.InvalidateOptions) error {
	err := c.store.Invalidate(options)

	if err == nil {
		c.stats.InvalidateSuccess++
	} else {
		c.stats.InvalidateError++
	}

	return err
}

// GetStore returns the store associated to this codec
func (c *Codec) GetStore() store.StoreInterface {
	return c.store
}

// GetStats returns some statistics about the current codec
func (c *Codec) GetStats() *Stats {
	return c.stats
}
