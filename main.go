package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func staticHandler(path string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		http.ServeFile(w, r, path)
		duration := time.Now().Sub(start)
		log.Printf("%s => static file: %s in %s ms", r.URL.Path, path, duration/time.Millisecond)
	}
}

func main() {
	fmt.Println("Welcome to the  Widowmaker Invitatinal")
	fmt.Printf("DB location: %s\n", "localhost:8888")
	http.HandleFunc("/", staticHandler("static/index.html"))
	http.ListenAndServe(":8080", nil)
}
