package handlers

import (
	"net/http"
)

// ServeCSS serves the static CSS file
func ServeCSS(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "frontend/styles.css")
}
