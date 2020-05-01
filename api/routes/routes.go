package routes

import (
	"fmt"
	"net/http"
	"time"
)

func GetTime(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format(time.RFC850)
	fmt.Fprintf(w, currentTime)
}

func GetToto(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "toto")
}

func GetRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "root")
}
