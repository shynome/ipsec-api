package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "7070"
	}
	addr := fmt.Sprintf("0.0.0.0:%d", port)
	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}

}
