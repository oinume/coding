package leetcode

import (
	"fmt"
	"io"
	"math"
	"strings"
	"sync"
	"testing"
)

/*
Design and implement a data structure for [Least Recently Used (LRU) cache](https://en.wikipedia.org/wiki/Cache_replacement_policies#LRU). It should support the following operations: get and put.

get(key) - Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.
put(key, value) - Set or insert the value if the key is not already present. When the cache reached its capacity, it should invalidate the least recently used item before inserting a new item.

The cache is initialized with a positive capacity.

Follow up:
Could you do both operations in O(1) time complexity?

Example:
```
LRUCache cache = new LRUCache(2) // capacity

cache.put(1, 1)
cache.put(2, 2)
cache.get(1)       // returns 1
cache.put(3, 3)    // evicts key 2
cache.get(2)       // returns -1 (not found)
cache.put(4, 4)    // evicts key 1
cache.get(1)       // returns -1 (not found)
cache.get(3)       // returns 3
cache.get(4)       // returns 4
```
*/
type LRUCache struct {
	capacity   int
	values     map[int]*entry
	currentAge int
	mutex      *sync.Mutex
}

type entry struct {
	value int
	age   int
}

func NewLRUCache(capacity int) LRUCache {
	if capacity < 0 {
		panic("can't specify negative capacity.")
	}
	return LRUCache{
		capacity:   capacity,
		values:     make(map[int]*entry, capacity),
		currentAge: 0,
		mutex:      new(sync.Mutex),
	}
}

func (c *LRUCache) Get(key int) int {
	e, ok := c.values[key]
	if !ok {
		return -1
	}
	c.mutex.Lock()
	e.age = c.currentAge
	c.currentAge++
	c.mutex.Unlock()
	return e.value
}

func (c *LRUCache) Put(key int, value int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if e, ok := c.values[key]; ok {
		e.value = value
		e.age = c.currentAge
		c.currentAge++
	} else {
		if len(c.values) >= c.capacity {
			leastAge := math.MaxInt32
			leastAgeKey := 0
			for key, e := range c.values {
				if e.age < leastAge {
					leastAge = e.age
					leastAgeKey = key
				}
			}
			if leastAgeKey != 0 {
				delete(c.values, leastAgeKey)
			}
		}
		c.values[key] = &entry{
			value: value,
			age:   c.currentAge,
		}
		c.currentAge++
	}
}

func (c *LRUCache) Dump(w io.Writer) error {
	t := `
{
  capacity: %v
  currentAge: %v
  values: { %v }
}
`
	a := make([]string, 0, len(c.values))
	for k, e := range c.values {
		a = append(a, fmt.Sprintf(`"%v":%+v`, k, e))
	}
	values := "{" + strings.Join(a, ", ") + "}"
	if _, err := fmt.Fprintf(w, t, c.capacity, c.currentAge, values); err != nil {
		return err
	}
	return nil
}

//LRUCache cache = new LRUCache( 2 /* capacity */ );
//
//cache.put(1, 1);
//cache.put(2, 2);
//cache.get(1);       // returns 1
//cache.put(3, 3);    // evicts key 2
//cache.get(2);       // returns -1 (not found)
//cache.put(4, 4);    // evicts key 1
//cache.get(1);       // returns -1 (not found)
//cache.get(3);       // returns 3
//cache.get(4);       // returns 4

func TestLRUCache_PutAndGet(t *testing.T) {
	cache := NewLRUCache(2)
	cache.Put(1, 1)
	cache.Put(2, 2)

	if got, want := cache.Get(1), 1; got != want {
		t.Errorf("got %v but want %v", got, want)
	}

	cache.Put(3, 3) // evicts key 2
	if got, want := cache.Get(2), -1; got != want {
		t.Errorf("got %v but want %v", got, want)
	}

	cache.Put(4, 4) // evicts key 1
	if got, want := cache.Get(1), -1; got != want {
		t.Errorf("got %v but want %v", got, want)
	}

	if got, want := cache.Get(3), 3; got != want {
		t.Errorf("got %v but want %v", got, want)
	}

	if got, want := cache.Get(4), 4; got != want {
		t.Errorf("got %v but want %v", got, want)
	}
}

func TestLRUCache_PutAndGet2(t *testing.T) {
	/*
		NewLRUCache(2)
		put(2, 1)
		put(2, 2)
		get(2)
		put(1, 1)
		put(4, 1)
		get(2)
	*/
	cache := NewLRUCache(2)
	cache.Put(2, 1)
	cache.Put(2, 2)
	if got, want := cache.Get(2), 2; got != want {
		t.Errorf("got %v but want %v", got, want)
	}

	cache.Put(1, 1)
	cache.Put(4, 1)
	if got, want := cache.Get(2), -1; got != want {
		t.Errorf("got %v but want %v", got, want)
	}
}
