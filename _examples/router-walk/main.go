package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/vksssd/bastiGO"
)

func main() {
	r := bastigo.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root."))
	})

	r.Route("/road", func(r bastigo.Router) {
		r.Get("/left", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("left road"))
		})
		r.Post("/right", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("right road"))
		})
	})

	r.Put("/ping", Ping)

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		route = strings.Replace(route, "/*/", "/", -1)
		fmt.Printf("%s %s\n", method, route)
		return nil
	}

	if err := bastigo.Walk(r, walkFunc); err != nil {
		fmt.Printf("Logging err: %s\n", err.Error())
	}
}

// Ping returns pong
func Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}
