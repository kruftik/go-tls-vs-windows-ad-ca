package main

import (
	// "fmt"
	// "io"
	"flag"
	"log"
	"net/http"
)

var (
	certFile = flag.String("cert", "server.crt", "A PEM eoncoded certificate file.")
	keyFile  = flag.String("key", "server.key", "A PEM encoded private key file.")
	caFile   = flag.String("CA", "ca.crt", "A PEM eoncoded CA's certificate file.")
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server.\n"))
}

func main() {
	flag.Parse()

	http.HandleFunc("/hello", HelloServer)
	err := http.ListenAndServeTLS(":443", *certFile, *keyFile, nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
