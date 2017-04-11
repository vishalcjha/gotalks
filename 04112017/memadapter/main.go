package main

import (
	"log"
	"net/http"

	"io/ioutil"
)

type memFunc func(w http.ResponseWriter, r []byte)

func main() {
	adaptedFunc := func(f memFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, req *http.Request) {
			//put sync.Pool and get bytes.Buffer and then use buf.ReadFrom(r.Body)
			body, _ := ioutil.ReadAll(req.Body)
			f(w, body)
		}
	}
	http.HandleFunc("/", adaptedFunc(func(w http.ResponseWriter, r []byte) {
		w.Write(r)
	}))
	log.Fatal(http.ListenAndServe("localhost:7777", nil))

}
