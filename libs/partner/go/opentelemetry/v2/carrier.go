package opentelemetryv2

import (
	"sync"

	"google.golang.org/grpc/metadata"
)

// Carrier TextMapCarrier provides a testing storage medium to for a
// TextMapPropagator. It records all the operations it performs.
type Carrier struct {
	mtx sync.Mutex

	gets []string
	sets [][2]string
	data map[string]string
}

func NewCarrier(m *metadata.MD) *Carrier {
	copied := make(map[string]string, m.Len())
	for k, v := range *m {
		if len(v) > 0 {
			copied[k] = v[0]
		}
	}
	return &Carrier{data: copied}
}

// Get returns the value associated with the passed key.
func (c *Carrier) Get(key string) string {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	c.gets = append(c.gets, key)
	return c.data[key]
}

// Set stores the key-value pair.
func (c *Carrier) Set(key, value string) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	c.sets = append(c.sets, [2]string{key, value})
	c.data[key] = value
}

// Keys Set stores the key-value pair.
func (c *Carrier) Keys() []string {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	i := 0
	keys := make([]string, len(c.data))
	for k := range c.data {
		keys[i] = k
		i++
	}
	return keys
}
