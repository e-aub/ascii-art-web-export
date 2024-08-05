package main

import (
	"fmt"
	"log"
	"net/http"

	"ascii-art-web/handlers"
	"ascii-art-web/utils"
)

func main() {
	utils.Banners = [5]string{"standard.txt", "shadow.txt", "thinkertoy.txt", "enigma.txt", "nirvana.txt"}
	// home page route, method : GET
	http.HandleFunc("/", handlers.HomeHandler)
	// ascii-art route, method : POST
	http.HandleFunc("/ascii-art", handlers.AsciiHandler)
	// Download file endpoint
	http.HandleFunc("/export", handlers.DownloadHandler)
	fmt.Println("starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
