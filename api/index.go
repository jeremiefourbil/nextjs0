package api

import (
	"api/routes"
	"github.com/gorilla/mux"
	"net/http"
)

var (
	router *mux.Router
)

func init() {
	// init router
	router = mux.NewRouter()

	apiRouter := router.PathPrefix("/api").Subrouter()

	// /api/(.*)
	apiRouter.HandleFunc("/time", routes.GetTime).Methods(http.MethodGet)
	apiRouter.HandleFunc("/toto", routes.GetToto).Methods(http.MethodGet)
	apiRouter.HandleFunc("/", routes.GetRoot).Methods(http.MethodGet)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}

func Sum(x int, y int) int {
	return x + y
}