package main

import (
	"fmt"
	"time"
)

func main() {

	// Bufffered and Unbuffered Channels

	unbufferedCh := make(chan int)  // Unbuffered channel
	bufferedCh := make(chan int, 2) // Buffered channel with capacity 2

	// Sending and receiving on unbuffered channel
	go func() {
		unbufferedCh <- 42 // Send value to unbuffered channel
	}()

	val := <-unbufferedCh // Receive value from unbuffered channel
	println("Received from unbuffered channel:", val)

	// Sending and receiving on buffered channel
	bufferedCh <- 1 // Send first value to buffered channel
	bufferedCh <- 2 // Send second value to buffered channel

	// Add one more value to demonstrate blocking behavior
	// This will block if the channel is full
	go func() {
		bufferedCh <- 3 // Send third value to buffered channel
		println("Sent third value to buffered channel")
	}()

	val1 := <-bufferedCh // Receive first value from buffered channel
	val2 := <-bufferedCh // Receive second value from buffered channel
	println("Received from buffered channel:", val1, val2)
	val3 := <-bufferedCh // Receive third value from buffered channel
	println("Received from buffered channel:", val3)

	// Closing Channels

	close(unbufferedCh) // Close unbuffered channel
	close(bufferedCh)   // Close buffered channel

	// Check if channels are closed
	if _, ok := <-unbufferedCh; !ok {
		println("Unbuffered channel is closed")
	}
	if _, ok := <-bufferedCh; !ok {
		println("Buffered channel is closed")
	}

	// // Goroutine Synchronization using channels
	done := make(chan bool)

	go func() {
		println("Goroutine is doing some work...")
		done <- true // Signal that work is done
	}()

	<-done // Wait for the goroutine to finish
	println("Goroutine has finished its work")

	// // Worker Pools / Concurrency Control
	// // Distribute work among multiple goroutines.
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	for w := 1; w <= 3; w++ {
		fmt.Println("Calling GORoutine with w=", w)
		go func(id int) {
			for j := range jobs {
				fmt.Printf("worker %d processing job %d\n", id, j)
				// Simulate work
				time.Sleep(time.Second)
				// Send result back
				results <- j * 2
			}
		}(w)
	}

	for j := 1; j <= 5; j++ {
		fmt.Printf("Sending job %d to jobs channel\n", j)
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= 5; a++ {
		// Wait for results
		fmt.Println("PRINT RESULT")
		fmt.Println(<-results)
	}
	close(results)

	// 	3. Pipeline Pattern
	// Chain multiple processing stages using channels.

	nums := make(chan int)
	squares := make(chan int)
	stop := make(chan struct{})

	go func() {
		for i := 1; i <= 5; i++ {
			// send numbers to nums channel
			nums <- i
		}
		close(nums)
	}()

	go func() {
		// read numbers from nums channel, square them, and send to squares channel
		for n := range nums {
			squares <- n * n
		}
		close(squares)
	}()

	// read squared values from squares channel
	for sq := range squares {
		fmt.Println("Square:", sq)
	}
	close(stop)

}
