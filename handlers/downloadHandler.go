package handlers

import (
	"bytes"
	"fmt"
	"net/http"

	asciiArt "ascii-art-web/ascii-art"
)

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	buffer := bytes.Buffer{}
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
	}
	textCookie, err := r.Cookie("text")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("bad request export"))
		return
	}
	bannerCookie, err := r.Cookie("banner")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("bad request export"))
		return
	}
	art, err := asciiArt.AsciiArt(textCookie.Value, bannerCookie.Value)
	// if an error occured while generating ascii art response with status 500 and a message
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}
	buffer.WriteString("<pre>")
	buffer.WriteString(art)
	buffer.WriteString("</pre>")
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	contentLen := buffer.Len()
	w.Header().Set("Content-Length", fmt.Sprintf("%d", contentLen))
	w.Header().Set("Content-Disposition", "attachment; filename=\"filename.html\"")
	w.Write(buffer.Bytes())
}
