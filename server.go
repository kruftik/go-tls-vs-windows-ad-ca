package main

import (
	"bufio"
	"crypto/tls"
	"flag"
	"log"
	"net"
)

var (
	certFile = flag.String("cert", "server.crt", "A PEM eoncoded certificate file.")
	keyFile  = flag.String("key", "server.key", "A PEM encoded private key file.")
	caFile   = flag.String("CA", "ca.crt", "A PEM eoncoded CA's certificate file.")
)

func main() {
	flag.Parse()

	log.SetFlags(log.Lshortfile)

	cer, err := tls.LoadX509KeyPair(*certFile, *keyFile)
	if err != nil {
		log.Println(err)
		return
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{cer},
	}

	ln, err := tls.Listen("tcp", ":443", config)
	if err != nil {
		log.Println(err)
		return
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	for {
		msg, err := r.ReadString('\n')
		if err != nil {
			log.Println(err)
			return
		}

		println(msg)

		n, err := conn.Write([]byte("world\n"))
		if err != nil {
			log.Println(n, err)
			return
		}
	}
}
