package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
)

func main() {
	port := "8080"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()
		fmt.Fprintf(w, "Hello from Go!\nHostname: %s\nPlatform: %s/%s\n", hostname, runtime.GOOS, runtime.GOARCH)
	})

	fmt.Printf("Listening on :%s\n", port)
	http.ListenAndServe(":"+port, nil)
}
