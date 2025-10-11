package main

import (
	"fmt"
	"net/http"
	"time"
)

// Job represents a URL to fetch
type Job struct {
	ID  int
	URL string
}

// Result represents the result of fetching a URL
type Result struct {
	JobID  int
	URL    string
	Status string
}

func worker(id int, jobs <-chan Job, results chan<- Result) {
	// simulate work
	fmt.Println("Waiting for jobs ䷄ with ID:", id)
	for job := range jobs {
		// simulate network delay
		time.Sleep(500 * time.Millisecond)

		resp, err := http.Get(job.URL)
		status := "Failed"
		if err == nil {
			status = resp.Status
			resp.Body.Close()
		}

		fmt.Printf("Worker %d fetched %s\n", id, job.URL)
		fmt.Println("Sending Results to Results Channel")
		results <- Result{JobID: job.ID, URL: job.URL, Status: status}
	}
}

func main() {
	jobs := make(chan Job, 10)
	results := make(chan Result, 10)

	// Start 3 worker goroutines
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Send jobs
	urls := []string{
		"https://google.com",
		"https://golang.org",
		"https://github.com",
		"https://medium.com",
		"https://vercel.com",
	}

	for i, url := range urls {
		jobs <- Job{ID: i + 1, URL: url}
	}
	close(jobs) // no more jobs to send

	// Receive results
	for i := 0; i < len(urls); i++ {
		res := <-results
		fmt.Printf("Result %d: %s (%s)\n", res.JobID, res.URL, res.Status)
	}
}
