package main

import (
	"log"
	"net/http"

	"github.com/YO-RO/kgb/handlers"
)

const port = ":8888"

func main() {
	log.Printf("Start Server - http://localhost%s\n", port)

	http.HandleFunc("/", handlers.HandleThreadView)
	http.HandleFunc("/post", handlers.HandleThreadPost)
	http.HandleFunc("/delete", handlers.HandleThreadDelete)
	log.Fatal(http.ListenAndServe(port, nil))
}
