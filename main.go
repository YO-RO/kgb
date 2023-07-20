package main

import (
	"log"
	"net/http"

	"github.com/YO-RO/kgb/handlers"
)

const port = ":8888"

func main() {
	log.Printf("Start Server - http://localhost%s\n", port)

	http.HandleFunc("/", handlers.ThreadsViewHandler)
	http.HandleFunc("/post", handlers.PostThreadHandler)
	http.HandleFunc("/delete", handlers.DeleteThreadHandler)
	log.Fatal(http.ListenAndServe(port, nil))
}
