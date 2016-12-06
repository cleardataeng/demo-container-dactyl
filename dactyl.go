package main

import "os"
import "net/http"
import "encoding/json"
import "log"

func slashHandler(w http.ResponseWriter, r *http.Request) {
	type DactylInfo struct {
		Hostname   string
		RemoteAddr string
	}

	name, err := os.Hostname()
	if err != nil {
		name = "(failed to load hostname)"
	}

	info := DactylInfo{
		name,
		r.RemoteAddr,
	}

	js, err := json.Marshal(info)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func logHandler(handler http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
	handler.ServeHTTP(w, r)
    })
}

func main() {
	http.HandleFunc("/", slashHandler)
	http.ListenAndServe(":8080", logHandler(http.DefaultServeMux))
}
