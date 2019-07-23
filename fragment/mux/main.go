package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/1", FirstHandler)
	r.HandleFunc("/2", SecondHandler)

	srv := &http.Server{
		Addr: "0.0.0.0:8080",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	srv.ListenAndServe()
}

func FirstHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Add("content-type", "application/json")
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	fmt.Println(1)
	fmt.Println(r.RemoteAddr)
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	io.WriteString(w, `~~~`)
}

func SecondHandler(w http.ResponseWriter, r *http.Request) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	s := buf.String()
	io.WriteString(w, s)
}
