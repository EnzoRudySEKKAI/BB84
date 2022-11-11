package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
)

const (
	ServerHost = "localhost"
	ServerPort = "8000"
	ServerType = "tcp"
)

func Client() {
	//establish connection
	connection, err := net.Dial(ServerType, ServerHost+":"+ServerPort)

	if err != nil {
		panic(err)
	}
	fmt.Println("test")
	alice := User{make([]qbit, 100)}

	for i := 0; i < 100; i++ {
		r := rand.Intn(2)
		value := rand.Intn(2)
		alice.getQbits()[i] = qbit{value, r}
	}

	toSend, _ := json.Marshal(alice.getQbits())
	fmt.Println("Sending to server...", string(toSend))
	///send some data
	_, err = connection.Write(toSend)

	buffer := make([]byte, 1024)
	_, err = connection.Read(buffer)
	var toread []qbit
	err = json.Unmarshal(buffer, &toread)
	cpt := 0
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	for i := 0; i < 100; i++ {
		if toread[i].getValue() == alice.getQbits()[i].getValue() {
			cpt++
		}
	}
	fmt.Println("Réussite: ", cpt, "%")
	defer func(connection net.Conn) {
		err := connection.Close()
		if err != nil {
			fmt.Println("Error closing:", err.Error())
		}
	}(connection)
}
