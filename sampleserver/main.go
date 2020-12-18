package helloworld

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

func handleHelloWorld(w http.ResponseWriter, r *http.Request) {
	urlPathElements := strings.Split(r.URL.Path, "/")
	text, ok := helloworld.Data[urlPathElements[1]]
	if ok {
		fmt.Fprintf(w, text)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - language not found!"))
	}
}

func main() {
	s := &http.Server{
		Addr:           ":8000",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
