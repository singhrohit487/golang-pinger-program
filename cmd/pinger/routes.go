package main

import "net/http"

// createMux returns the multiplexer we use to route our requests
func createMux() *http.ServeMux {
	mux := http.NewServeMux()

	// root endpoint - say hi!
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello world"))
	})

	// liveness check
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("alive"))
	})

	// readiness check
	mux.HandleFunc("/readyz", func(w http.ResponseWriter, r *http.Request) {
		if readiness["target_up"] {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("ready"))
		} else {
			w.WriteHeader(http.StatusTeapot)
			w.Write([]byte("not ready"))
		}
	})
	return mux
}
