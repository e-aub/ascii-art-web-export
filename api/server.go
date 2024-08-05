package main

import (
	"fmt"
	"log"
	"net/http"

	"api/utils"

	"api/handlers"
)

func main() {
	utils.Banners = [5]string{"standard.txt", "shadow.txt", "thinkertoy.txt", "enigma.txt", "nirvana.txt"}
	http.HandleFunc("/", handlers.AsciiArtServer)
	fmt.Println("starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
