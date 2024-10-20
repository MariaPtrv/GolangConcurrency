package concurMapBuffChannelv2

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

type concurrentMap struct {
	Map   map[string]int
	mutex chan int
}

func (c concurrentMap) get(key string) (int, bool) {
	c.mutex <- 1
	v, has := c.Map[key]
	<-c.mutex
	return v, has
}

func (c concurrentMap) set(key string, val int) {
	c.mutex <- 1
	c.Map[key] = val
	<-c.mutex
}

func ConcurMapBuffChannel() {
	var wg sync.WaitGroup

	m := concurrentMap{
		Map:   make(map[string]int),
		mutex: make(chan int, 1),
	}

	m.set("one", 1)
	v, _ := m.get("one")
	fmt.Println(v)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i < 10; i++ {
			key := strings.Join([]string{"key", strconv.Itoa(i)}, " ")

			//fmt.Println("Lock write")
			m.set(key, i)

		}
	}()

	wg.Add(1)
	go func() {
		t := time.NewTicker(time.Second * 2)
		<-t.C

		defer wg.Done()
		for i := 1; i < 10; i++ {
			key := strings.Join([]string{"key", strconv.Itoa(i)}, " ")

			//fmt.Println("Lock read")
			v, _ := m.get(key)

			fmt.Println(v)
		}
	}()

	wg.Wait()
}
