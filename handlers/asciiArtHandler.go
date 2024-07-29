package handlers

import (
	"net/http"

	asciiArt "ascii-art-web/ascii-art"
	"ascii-art-web/utils"
)

// ascii-art controller that recives a post request with necessary data and generates ascii art
func AsciiHandler(w http.ResponseWriter, r *http.Request) {
	// check if method is post
	if r.Method == http.MethodPost {
		// extract form data from the request
		form := utils.FormValues{
			Text:   r.FormValue("text"),
			Banner: r.FormValue("banner"),
		}
		// check if the banner is valid
		if utils.IsBanner(form.Banner) {
			var err error
			// generate ascii art
			utils.Art, err = asciiArt.AsciiArt(form.Text, form.Banner)
			// if an error occured while generating ascii art response with status 500 and a message
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Internal server error"))
				return
			}
			// if the art is generated succusfully redirect to home page
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		// respond with status bad request and a message if the banner is invalid
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid Banner"))
		return
	}
	// Respond with status 405 and a message to inform that the method used is not allowed
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("Method not allowed"))
}
