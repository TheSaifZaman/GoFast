package router

import (
	"fmt"
	"github.com/TheSaifZaman/GoFast/config"
	"net/http"
)

func NewRouter() http.Handler {
	mainMux := http.NewServeMux()
	mainMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to Go-Fast! (top-level)")
	})
	config.ConfigureAPIs(mainMux)
	return globalMiddleware(mainMux)
}

func globalMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("[GLOBAL MIDDLEWARE] Request received:", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
