package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
)

func Client() {
	fmt.Println("start client")
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("Connection error", err)
	}
	encoder := gob.NewEncoder(conn)
	p := &Qbit{1, 0}
	encoder.Encode(p)
	conn.Close()
	fmt.Println("done")
}
