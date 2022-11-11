package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "8000"
	SERVER_TYPE = "tcp"
)

func Server() {
	fmt.Println("Server Running...")
	server, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer func(server net.Listener) {
		err := server.Close()
		if err != nil {
			fmt.Println("Error closing:", err.Error())
		}
	}(server)
	fmt.Println("Listening on " + SERVER_HOST + ":" + SERVER_PORT)
	fmt.Println("Waiting for client...")
	for {
		connection, err := server.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		fmt.Println("client connected")
		go processClient(connection)
	}
}
func processClient(connection net.Conn) {
	buffer := make([]byte, 1024)
	_, err := connection.Read(buffer)

	var toread []qbit
	err = json.Unmarshal(buffer, &toread)
	if err != nil {
		return
	}
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}

	bob := User{make([]qbit, len(toread))}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(toread); i++ {
		basis := rand.Intn(2)
		bob.getQbits()[i] = qbit{toread[i].measure(basis), toread[i].getPolarisation()}
	}
	toSend, _ := json.Marshal(bob.getQbits())
	fmt.Println(toSend)
	_, err = connection.Write(toSend)
	err = connection.Close()
	if err != nil {
		return
	}
}
