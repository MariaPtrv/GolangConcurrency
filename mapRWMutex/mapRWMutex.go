package mapRWMutex

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

type concurrentMap struct {
	Map     map[string]int
	rwMutex sync.RWMutex
}

func (c *concurrentMap) get(key string) (int, bool) {
	c.rwMutex.RLock()
	defer c.rwMutex.RUnlock()
	v, has := c.Map[key]

	return v, has
}

func (c *concurrentMap) set(key string, val int) {
	c.rwMutex.Lock()
	defer c.rwMutex.Unlock()
	c.Map[key] = val

}

func MapRWMutex() {
	var wg sync.WaitGroup

	m := &concurrentMap{
		Map: make(map[string]int),
	}

	m.set("one", 1)
	v, _ := m.get("one")
	//_, _ = m.get("one")
	fmt.Println(v)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i < 10; i++ {
			key := strings.Join([]string{"key", strconv.Itoa(i)}, " ")
			m.set(key, i)

		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i < 10; i++ {
			key := strings.Join([]string{"key", strconv.Itoa(i)}, " ")
			v, _ := m.get(key)

			fmt.Println(v)
		}
	}()

	wg.Wait()
}
