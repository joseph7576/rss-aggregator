package main

import "net/http"

// special function signature for handling request - go std lib.
func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, struct{}{})
}
