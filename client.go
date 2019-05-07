package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"log"

	"io/ioutil"
)

var (
	certFile = flag.String("cert", "server.crt", "A PEM eoncoded certificate file.")
	keyFile  = flag.String("key", "server.key", "A PEM encoded private key file.")
	caFile   = flag.String("CA", "ca.crt", "A PEM eoncoded CA's certificate file.")

	serverHostPort = flag.String("connect", "127.0.0.1:443", "Server host:port to connect to")
)

func main() {
	flag.Parse()

	log.SetFlags(log.Lshortfile)

	caCert, err := ioutil.ReadFile(*caFile)
	if err != nil {
		log.Fatal(err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	conf := &tls.Config{
		RootCAs: caCertPool,
		//InsecureSkipVerify: true,
	}

	conn, err := tls.Dial("tcp", *serverHostPort, conf)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	n, err := conn.Write([]byte("hello\n"))
	if err != nil {
		log.Println(n, err)
		return
	}

	buf := make([]byte, 100)
	n, err = conn.Read(buf)
	if err != nil {
		log.Println(n, err)
		return
	}

	println(string(buf[:n]))
}
