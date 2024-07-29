package utils

import "net/http"

// type to hold form values
type FormValues struct {
	Text   string
	Banner string
}

// Art variable that will hold ascii art and banners array to validate banners
var (
	Art     string
	Banners [5]string
)

// function to check if banner matches one of the banners in banners array
func IsBanner(item string) bool {
	for _, banner := range Banners {
		if banner == item {
			return true
		}
	}
	return false
}

// not found controller that informs client in a response with 404 status and a message if a ressource is not found
func NotFound(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(message))
}
