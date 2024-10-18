package main

import (
	"concurrency/rangeChannelSelect"
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")

	//wg := sync.WaitGroup{}
	//wg.Add(1)
	//signalChannel.SignalChannel()
	// signalNotify.SignalNotify()
	//contextPanic.ContextPanic()
	// go func() {
	// 	defer wg.Done()
	// 	concurMapError.ConcurMapError()
	// }()

	//wg.Wait()

	//concurMapMutex.ConcurMapMutex()
	//concurMapBuffChannel.ConcurMapBuffChannel()

	//concurMapBuffChannelv2.ConcurMapBuffChannel()
	//mapRWMutex.MapRWMutex()
	rangeChannelSelect.RangeChannelSelect()
}
