package main

import (
	"fmt"
	"log"
	"net"
)

type GeoPoint struct {
	a       int     // simple int
	aLong   int64   // long int
	aDouble float64 // double float64
}

func (pt *GeoPoint) printStats() {
	fmt.Println(pt.a, pt.aLong, pt.aDouble)
}

func main() {

	ln, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Listening on ", ln.Addr())

	connection, err := ln.Accept()
	if err == nil {
		// fmt.Println("Connection: ", connection)
		go serve(connection)
	}
}

func serve(c net.Conn) {
	fmt.Fprintf(c, "Hello there %v\n", c.RemoteAddr())
}
