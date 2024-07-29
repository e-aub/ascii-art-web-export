package handlers

import (
	"ascii-art-web/utils"
	"html/template"
	"net/http"
)

// home controller to inject ascii art in the template and parse it to serve it to the client
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	//if no route is matched mux will choose /, Since / is a catch-all route, it handles /unknown,
	// Because / is the least specific match, it acts as a fallback for any request that doesn't match a more specific route.
	// Thats why i check here if the route is exactly / if not the server responds with 404 status and a not found message
	if r.URL.Path != "/" {
		utils.NotFound(w, "404 page not found")
		return
	}
	//Verifying that the method is get
	if r.Method == http.MethodGet {
		//Parse the html template
		tmpl, err := template.ParseFiles("templates/index.html")
		//if an error occured while parsing the server responds with status 500 and internal server error message
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal server errorf"))
			return
		}
		//respond with status 200 OK and serve the template with ascii art generated before when a post request recived by /ascii-art
		w.WriteHeader(http.StatusOK)
		tmpl.Execute(w, utils.Art)
		//erase Art variable for next use
		utils.Art = ""
		return
	}
	//Respond with status 405 and a message to inform that the method used is not allowed
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("Method not allowed"))

}
