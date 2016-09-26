package main

import "os"
import "net/http"
import "encoding/json"

func handler(w http.ResponseWriter, r *http.Request) {
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

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
