package config

import (
	"fmt"
	"net/http"
)

func ConfigureAPIs(mainMux *http.ServeMux) {
	v1Mux := http.NewServeMux()
	mainMux.Handle("/v1/", http.StripPrefix("/v1", v1Mux))
	v2Mux := http.NewServeMux()
	mainMux.Handle("/v2/", http.StripPrefix("/v2", v2Mux))

	v1Mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from v1!")
	})
	v2Mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from v2!")
	})
}

func apiResource(mux *http.ServeMux, basePath string) {
	mux.HandleFunc(basePath, func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			fmt.Fprintf(w, "Index of %s\n", basePath)
		case http.MethodPost:
			fmt.Fprintf(w, "Create new %s\n", basePath)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc(basePath+"/", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Path[len(basePath)+1:]
		if id == "" {
			http.Error(w, "Missing resource ID", http.StatusBadRequest)
			return
		}

		switch r.Method {
		case http.MethodGet:
			fmt.Fprintf(w, "Show %s with ID %s\n", basePath, id)
		case http.MethodPut:
			fmt.Fprintf(w, "Update %s with ID %s\n", basePath, id)
		case http.MethodDelete:
			fmt.Fprintf(w, "Delete %s with ID %s\n", basePath, id)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
}
