// Channels are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and receive those values into another goroutine.
package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a new channel with make(chan val-type). Channels are typed by the values they convey.
	messages := make(chan string)

	// Send a value into a channel using the channel <- syntax. Here we send "ping" to the messages channel we made above, from a new goroutine.
	go func() { messages <- "ping" }()

	// The <-channel syntax receives a value from the channel. Here we’ll receive the "ping" message we sent above and print it out.
	msg := <-messages
	fmt.Println(msg)

	// Test
	//go func() {
	//	for {
	//		fmt.Println("1")
	//		v, ok := <-messages
	//		fmt.Println(ok)
	//		if !ok {
	//			break
	//		} // use v
	//		fmt.Println(v)
	//	}
	//}()
	//go func() { messages <- "pong2" }()
	//fmt.Scanln()
	// Test End
	// When we run the program the "ping" message is successfully passed from one goroutine to another via our channel.

	// By default sends and receives block until both the sender and receiver are ready.
	// This property allowed us to wait at the end of our program for the "ping" message without having to use any other synchronization.

	/**
	By default channels are unbuffered,
	meaning that they will only accept sends (chan <-) if there is a corresponding receive (<- chan) ready to receive the sent value.
	Buffered channels accept a limited number of values without a corresponding receiver for those values.
	*/
	// Here we make a channel of strings buffering up to 2 values.
	m := make(chan string, 2)

	// Because this channel is buffered, we can send these values into the channel without a corresponding concurrent receive.
	m <- "buffered"
	m <- "channel"

	// Later we can receive these two values as usual.
	fmt.Println(<-m)
	fmt.Println(<-m)

	/**
	We can use channels to synchronize execution across goroutines. Here’s an example of using a blocking receive to wait for a goroutine to finish.
	*/

	// Start a worker goroutine, giving it the channel to notify on.
	done := make(chan bool, 1)
	go worker(done)
	// Block until we receive a notification from the worker on the channel.
	<-done

	/**
	When using channels as function parameters,
	you can specify if a channel is meant to only send or receive values.
	This specificity increases the type-safety of the program.
	*/
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)

}

// This is the function we’ll run in a goroutine. The done channel will be used to notify another goroutine that this function’s work is done.
func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	// Send a value to notify that we’re done.
	done <- true
}

// This ping function only accepts a channel for sending values. It would be a compile-time error to try to receive on this channel.
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// The pong function accepts one channel for receives (pings) and a second for sends (pongs).
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
