package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola desde un unikernel con Ops 🚀")
}

func main() {
	start := time.Now()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK")
	})

	go func() {
		http.ListenAndServe(":8080", nil)
	}()

	fmt.Println("Server ready in:", time.Since(start))
	select {}
}