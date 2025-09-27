/*
Key Concepts
Concurrency: The program uses goroutines to perform HTTP requests concurrently.

	This allows it to handle multiple URLs simultaneously, improving efficiency.

Context: The context package is used to manage the timeout for the requests.

	If any request takes longer than 5 seconds, it will be cancelled.

Channels: Channels are used to collect results from the concurrent goroutines.

	This allows the main function to print the results as they are completed.
*/
package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	executeCancelTaskWithContext()
	propogatingContextWithValues()
	concurrentOperations()
	// this function has premature termination so use it at the end
	// Anything after this will not be executed
	performConcurrentTasks()
}

func concurrentOperations() {
	// Creating a context with a 5-second timeout
	// If the operations don't complete within this time, they will be cancelled.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// ensures that the context is cancelled when the main function completes, releasing any resources associated with it.
	defer cancel() // Ensure the context is cancelled to free resources

	// List of URLs to fetch
	urls := []string{
		"https://jsonplaceholder.typicode.com/todos/1",
		"https://jsonplaceholder.typicode.com/posts/1",
		"https://jsonplaceholder.typicode.com/users/1",
	}

	/* Channel to receive the results
	creates a channel to receive the results
	from the concurrent fetch operations. */
	results := make(chan string)

	/* Launch a goroutine for each URL
	loop iterates over the URLs, launching a goroutine for each
	one with the fetchAPI function.
	This enables concurrent execution of the HTTP requests. */
	for _, url := range urls {
		go fetchAPI(ctx, url, results)
	}

	/* Collect results from all goroutines
	loop waits for results from the results channel.
	It prints each result to the console as they come in. */
	for range urls {
		fmt.Println(<-results)
	}
}
func fetchAPI(ctx context.Context, url string, results chan<- string) {
	// creates a new HTTP GET request with the provided context.
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	/* If there is an error creating the request,
	it sends an error message to the results channel and returns. */
	if err != nil {
		results <- fmt.Sprintf("Error creating request for %s: %s", url, err.Error())
		return
	}

	// uses the default HTTP client.
	client := http.DefaultClient
	// sends the HTTP request.
	resp, err := client.Do(req)
	if err != nil {
		// If an error occurs, it sends an error message to the results channel and returns.
		results <- fmt.Sprintf("Error making request to %s: %s", url, err.Error())
		return
	}
	/* ensures that the response body is closed after
	it has been read to free resources. */
	defer resp.Body.Close()

	/* sends a formatted string containing the URL and its HTTP status code
	to the results channel. */
	results <- fmt.Sprintf("Response from %s: %d", url, resp.StatusCode)
}

func performConcurrentTasks() {
	// Creating a context with a 2-second timeouts
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go performTask(ctx)

	// In this example, the performTask function simulates a long-running task that takes 5 seconds to complete. However, since the context has a timeout of only 2 seconds, the operation is terminated prematurely, resulting in a timeout.
	select {
	case <-ctx.Done():
		fmt.Println("Task timed out intentionally to demonstrate the option of context with timeout: [", ctx.Err(), "]")
	}
}

func performTask(ctx context.Context) {
	select {
	// Actual task takes 5 seconds to complete
	case <-time.After(5 * time.Second):
		fmt.Println("Task completed successfully")
	}
	_ = ctx.Value("")
}

func propogatingContextWithValues() {

	ctx := context.Background()

	ctx = context.WithValue(ctx, "UserID", 123)

	go RetrieveUserFromContextValue(ctx)

}

func RetrieveUserFromContextValue(ctx context.Context) {
	// Retrieve the value from the context
	userID := ctx.Value("UserID")
	fmt.Println("User ID:", userID)
	// Retrieve the value from the context and convert it to an integer
	fmt.Println("Processing request for User ID:", userID.(int))
}

func executeCancelTaskWithContext() {
	ctx, cancel := context.WithCancel(context.Background())

	go PerformCancelTaskWithContext(ctx)

	time.Sleep(2 * time.Second)
	cancel()

	time.Sleep(1 * time.Second)
}
func PerformCancelTaskWithContext(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Task cancelled")
			return
		default:
			// Perform task operation
			fmt.Println("Performing Cancellable task...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func ExecuteTaskWithDeadline() {

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
	defer cancel()

	go performTask(ctx)

	time.Sleep(3 * time.Second)

}
func PerformTaskWithDeadline(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Task completed or deadline exceeded:", ctx.Err())
		return
	}
}
