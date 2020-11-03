package main

import (
    "context"
    "github.com/gorilla/mux"
    "log"
    "net/http"
    "os"
    "os/signal"
    "time"
)

var bindAddress = env.string("BIND_ADDRESS",false, ":9090", "Bind address for the server")

func main() {
     
     env.Parse()

     l :=log.new(os.stdout, "products-api",log.Lstdflags)

     // create the handlers
     ph := handlers.NewProducts(l)

     // create a new serve mux and register the handlers
     sm :=mux.newRouter()
    
    getRouter := sm.Methods(http.MethodGet).Subrouter()
    getRouter.HandleFunc("/",ph.GetProducts)

    putRouter := sm.Methods(http.MethodPut).Subrouter()
    putRouter.HandleFunc("/{id:[0-9]+}", ph.updateProducts)
	putRouter.Use(ph.MiddlewareValidateProduct)
	
    postrouter := sm.Methods(http.methodpost).Subrouter()
	postrouter.HandleFunc("/", ph.AddProduct)
	postRouter.Use(ph.MiddlewareValidateProduct)

     //sm.Handle("/products", ph)

     // craete a new server
     s := http.Server{
         Addr:     *bindAddress,     // configure the bind address
         Handler:   sm,              // set the default handler
         ErrorLog:  l,               // set the logger for the server
         ReadTimeout: 5 * time.Second, // max time to read request from the client
         WriteTimeout: 10 * time second, // max time to write response to the client
         Idletimeout: 120 * time.Second, // max time for connections using TCP Keep-Alive
     
	 }

	 
	}
