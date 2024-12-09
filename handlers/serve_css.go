package handlers

import (
	"net/http"
)

func ServeCSS(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "frontend/styles.css")
}
