package main

import (
	"log"
	"net/http"

	"github.com/MediaMath/playground/adaptersetup"
)

type errorFunc func(w http.ResponseWriter, req *http.Request) error

func main() {
	adaptedFunc := func(f errorFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, req *http.Request) {
			if err := f(w, req); err != nil {
				w.Write([]byte(err.Error()))
			} else {
				w.Write([]byte("Hello World"))
			}
		}
	}
	http.HandleFunc("/", adaptedFunc(adaptersetup.HelloWithError))
	log.Fatal(http.ListenAndServe("localhost:7777", nil))

}
