package contextPanic

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func ContextPanic() {
	ctx, cancel := context.WithCancel(context.TODO())
	wg := sync.WaitGroup{}

	d1 := time.Now().Add(2 * time.Second)
	ctx1, cancel1 := context.WithDeadline(ctx, d1)

	wg.Add(1)
	go One(ctx1, &wg)

	d2 := time.Now().Add(4 * time.Second)
	ctx2, cancel2 := context.WithDeadline(ctx, d2)

	wg.Add(1)
	go Two(ctx2, &wg)

	cancel()
	fmt.Println("Cancel")
	wg.Wait()
	// time.Sleep(time.Second * 4)

	cancel2()
	fmt.Println("Cancel2")
	cancel1()
	fmt.Println("Cancel1")

}

func One(ctx context.Context, wg *sync.WaitGroup) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
			wg.Done()
		}
	}()
	panic("panic")
	select {
	case <-ctx.Done():
		fmt.Println("Done 1")
	}

	wg.Done()

}

func Two(ctx context.Context, wg *sync.WaitGroup) {

	select {
	case <-ctx.Done():
		fmt.Println("Done 2")
	}

	wg.Done()
}
