package main

import (
	"net/http"
	"log"
	"io/ioutil"
)

import log

func main() {
  http.Handlefunc("/", func(http.ResponseWriter, *http.Request) {
	  log.Println("Hello World")
	  d, err := ioutl.readAll(r.Body)
	   if err != nil {
		   http.Error(rw, "Oops", http.StatusBadRequest)
		   return
	   }
	  
	  fmt.Fprintf(rw, "Hello %s", d)

  })

  http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
	  log.println("Goodbye World")
  })

  http.ListenAndServe(":9090", nil)
}