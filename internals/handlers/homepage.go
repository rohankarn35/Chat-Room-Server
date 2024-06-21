package handlers

import (
	"fmt"
	"net/http"
)

func Homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to Chat Server")
}
