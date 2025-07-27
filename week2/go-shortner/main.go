package main

import (
	"fmt"
	"goLabs/go-shortner/internal/handler"
	"goLabs/go-shortner/internal/store"
	"net/http"
)

func main(){

	m := &store.MemoryStore{
		Urls: make(map[string]string),
		Id: 0,
	}

	mux := http.NewServeMux()

	mux.Handle("POST /shorten", handler.ShortenHandler(m))
	mux.Handle("GET /{short_code}", handler.RedirectHandler(m))

	fmt.Println("Starting Server At 8080")
	http.ListenAndServe(":8080",mux)
}
