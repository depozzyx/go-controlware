package server

import (
	"log"
	"net/http"

	"controlware/shared"

	"github.com/gorilla/mux"
)

var (
	commands  []shared.Command = make([]shared.Command, 0)
	lastClear int64            = shared.LastTimestamp()
)

func Run(host string) {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/commands/get", get)
	r.HandleFunc("/commands/add", add).Queries("pwd", "dpx")
	r.HandleFunc("/commands/clear", clear).Queries("pwd", "dpx")
	r.HandleFunc("/shebangs/output", output).Methods("POST")
	r.HandleFunc("/file/build", buildFile).Queries("file", "{file}").Methods("GET")
	r.HandleFunc("/file/output", outputFile).Queries("file", "{file}").Methods("GET")

	r.Use(logMiddleware(host))

	http.ListenAndServe(host, r)
}

func logMiddleware(host string) func(http.Handler) http.Handler {
	log.Printf("[NOTIFICATION] Server running on %s \n", host)
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
			log.Printf("[%s] path %s | ip %s \n", r.Method, r.URL.Path, r.RemoteAddr)
		})
	}
}
