package examples

import (
	"fmt"
	"sync"
	"time"
)

func Foo(id int) {
	time.Sleep(1 * time.Second)
	fmt.Println("Finished with id: ", id)

}

func SeqExample() {
	t1 := time.Now()
	for i := 0; i < 5; i++ {
		go Foo(i)
	}
	total := time.Since(t1)
	fmt.Println("Total time: ", total)
}

func ChannelsExample() {
	messages := make(chan string)
	go func() { messages <- "Zajonc1" }()
	go func() { messages <- "Zajonc2" }()
	time.Sleep(1 * time.Second)
	go func() { messages <- "Zajonc3000" }()

	messageFromChannel := <-messages
	// msg := <-messages // Error: Channel is not buffered
	fmt.Println("Message from channel: ", messageFromChannel)
	// fmt.Println("Message from channel: ", msg)
}

func ChannelsBufferingExample() {
	channelOne := make(chan string, 2)
	channelTwo := make(chan string, 2)
	go func() { channelOne <- "Zajonc4" }()
	go func() { channelTwo <- "Zajonc3" }()
	msg1 := <-channelOne
	msg2 := <-channelTwo
	fmt.Println("[Buffered] Message from channel: ", msg1)
	fmt.Println("[Buffered] Message from channel: ", msg2)
	close(channelOne)
	close(channelTwo)
}

func SelectExample() {
	channelOne := make(chan string)
	channelTwo := make(chan string)
	go func() { channelOne <- "Msg1" }()
	go func() { channelTwo <- "Msg2" }()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-channelOne:
			fmt.Println("Received", msg1)
		case msg2 := <-channelTwo:
			fmt.Println("Received", msg2)
		}
	}
}

func WaitGroupExample() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		fmt.Println("Hello")
		defer wg.Done()
	}()
	wg.Wait()
}
