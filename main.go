package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("/test response"))
	})

	http.ListenAndServe(":8000", mux)
}
