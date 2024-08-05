package handlers

import (
	"encoding/json"
	"net/http"

	asciiArt "api/ascii-art"
	"api/utils"
)

// ascii-art controller that recives a post request with necessary data and generates ascii art
func AsciiArtServer(w http.ResponseWriter, r *http.Request) {
	// check if method is post
	if r.Method == http.MethodPost {
		form := utils.FormValues{}
		// extract form data from the request
		err := json.NewDecoder(r.Body).Decode(&form)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("bad request"))
			return
		}
		// check if the banner is valid
		if utils.IsBanner(form.Banner) {
			var err error
			// generate ascii art
			utils.Data.Art, err = asciiArt.AsciiArt(form.Text, form.Banner)
			// if an error occured while generating ascii art response with status 500 and a message
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Internal server error"))
				return
			}
			// Set cookie
			textCookie := http.Cookie{Name: "text", Value: form.Text, Path: "/export", MaxAge: 3600}
			bannerCookie := http.Cookie{Name: "banner", Value: form.Banner, Path: "/export", MaxAge: 3600}
			http.SetCookie(w, &textCookie)
			http.SetCookie(w, &bannerCookie)
			// if the art is generated succusfully redirect to home page
			utils.Data.DownloadButton = true
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
