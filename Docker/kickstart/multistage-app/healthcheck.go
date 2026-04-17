package main

import (
	"net/http"
	"os"
	"time"
)

func main() {
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get("http://localhost:8080/")
	if err != nil {
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		os.Exit(1)
	}
}
