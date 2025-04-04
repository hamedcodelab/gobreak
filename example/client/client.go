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

	gbrk := gobreak.NewBreaker(gobreak.WithFailureThreshold(1), gobreak.WithRecoveryTime(time.Minute*4), gobreak.WithHalfOpenMaxRequests(2))

	var client http.Client
	url := "http://localhost:8080"

	gbrk.Execute(func() error {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
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

	gbrk.Execute(func() error {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
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

	time.Sleep(time.Second * 5)

	gbrk.Execute(func() error {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
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

	time.Sleep(time.Second * 5)

	err := gbrk.Execute(func() error {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
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
