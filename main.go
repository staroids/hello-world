package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	current_time := time.Now().Local()
	fmt.Fprintf(w, "Hello World! -  %s\n", current_time.Format("2006-01-02 15:04:05"))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":2222", nil)
}
