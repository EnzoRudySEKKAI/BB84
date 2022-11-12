package main

import (
	"encoding/gob"
	"fmt"
	"math/rand"
	"net"
	"time"
)

func handleConnection(conn net.Conn) {
	dec := gob.NewDecoder(conn)
	p := &Qbit{}
	dec.Decode(p)
	rand.Seed(time.Now().UTC().UnixNano())
	r := rand.Intn(2)
	qbitDecoded := Qbit{p.measure(r), p.getPolarisation()}
	fmt.Printf("Received : %+v\n", qbitDecoded)
	conn.Close()
}

func Server() {
	fmt.Println("start")
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
	}
	for {
		conn, err := ln.Accept() // this blocks until connection or error
		if err != nil {
			// handle error
			continue
		}
		go handleConnection(conn) // a goroutine handles conn so that the loop can accept other connections
	}
}
