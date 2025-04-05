package main

import (
	"context"
	"fmt"
	"github.com/hamedcodelab/gobreak"
	"log"
	"net/http"
	"time"
)

func main() {
	// Create a new breaker with failure threshold 1 and a long recovery time
	gbrk := gobreak.NewBreaker(
		gobreak.WithFailureThreshold(1),          // Threshold for failure count before breaking
		gobreak.WithRecoveryTime(time.Second*10), // Recovery time (open state duration)
		gobreak.WithHalfOpenMaxRequests(2),       // Max requests allowed while half-open
	)

	// Create an HTTP client
	var client http.Client
	// todo Change to your actual endpoint
	url := "http://localhost:8080"

	// Simulate a failure to trip the breaker
	fmt.Println("Executing request 1 (failure expected)...")
	err := gbrk.Execute(func() error {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
		if err != nil {
			return err
		}
		_, err = client.Do(req)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		fmt.Println("Request failed, circuit breaker tripped")
	}

	// Wait for the breaker to open and then recover
	fmt.Println("Waiting for recovery time (10 seconds)...")
	time.Sleep(10 * time.Second)

	// Try a few requests while the breaker is in the half-open state
	for i := 0; i < 3; i++ {
		fmt.Printf("\nExecuting request %d (half-open state)...\n", i+1)
		err = gbrk.Execute(func() error {
			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			defer cancel()
			req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
			if err != nil {
				return err
			}
			_, err = client.Do(req)
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			fmt.Println("Request failed, still in half-open state")
		} else {
			fmt.Println("Request successful, breaker should move to closed state now")
			break // Break after a successful request in half-open state
		}
	}

	// Once the breaker has transitioned to closed, try again
	fmt.Println("\nWaiting for breaker to move to 'closed' state after success...")
	time.Sleep(2 * time.Second)

	// Check final status by making a successful request
	fmt.Println("\nExecuting final request (should succeed now)...")
	err = gbrk.Execute(func() error {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
		if err != nil {
			return err
		}
		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		fmt.Println("Response Status:", resp.Status)
		return nil
	})
	if err != nil {
		log.Fatalf("HTTP request failed: %v", err)
	}

	fmt.Println("HTTP request successful")
}
