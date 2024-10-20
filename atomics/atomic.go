package atomics

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func AddAtomic() {
	var counter int64

	wg := sync.WaitGroup{}

	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
			fmt.Println(i)
		}()
	}

	wg.Wait()
	fmt.Println("==========")
	fmt.Println(counter)
}

//LoadT
//StoreT
//SwapT

//CompareAndSwapT



func CompareAndSwapAtomic() {
	//операцию выполнит только одна горутина
	//sync.Once использует atomic.CompareAndSwap
	var counter int64

	wg := sync.WaitGroup{}

	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			if !atomic.CompareAndSwapInt64(&counter, 0, 1) {
				return 
			}
			fmt.Println(i)
		}()
	}

	wg.Wait()
	fmt.Println("==========")
	fmt.Println(counter)
}
