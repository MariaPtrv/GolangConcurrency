package concurMapError

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

type concurrentMap struct {
	Map map[string]int
}

func (c concurrentMap) get(key string) (int, bool) {
	v, has := c.Map[key]
	fmt.Printf("GET %s: %d\n", key, v)
	return v, has
}

func (c concurrentMap) set(key string, val int) {
	c.Map[key] = val
	fmt.Printf("SET %s: %d\n", key, val)
}

func ConcurMapError() {
	var wg sync.WaitGroup

	m := concurrentMap{
		Map: make(map[string]int),
	}

	m.set("one", 1)
	v, _ := m.get("one")
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
			key := strings.Join([]string{"key 2", strconv.Itoa(i)}, " ")
			m.set(key, i)
		}
	}()

	wg.Wait()
}
