package main

import (
	"fmt"
	"log"
	"net/http"
)

var Version string

func getVersion() string {
	if Version == "" {
		Version = "unknown"
	}
	return Version
}

const port = 8081

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("Running server version: %s\n", getVersion())))
	})

	fmt.Printf("Starting server on :%d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal("Unable to start server: ", err)
	}
}
