// Qbit bb84 protocol
package main

import (
	"fmt"
	"os"
)

func main() {
	nbQ := 100
	pola := 2
	nbUsers := 3
	if len(os.Args) > 1 {
		if os.Args[1] == "server" {
			Server()
		} else if os.Args[1] == "client" {
			Client()
		} else if os.Args[1] == "75" {
			test75(nbQ, pola)
		} else if os.Args[1] == "62" {
			test62(nbQ, pola)
		} else if os.Args[1] == "3users" {
			test3users(nbQ, pola)
		} else if os.Args[1] == "Nusers" {
			testNusers(nbQ, pola, nbUsers)
		}
	} else {
		fmt.Println("Liste des arguments :\n server : lancer le serveur\n client : lancer le client\n 75 : lancer le test de 75%\n 62 : lancer le test de 62%\n 3users : lancer le test de 3 utilisateurs\n Nusers : lancer le test de N utilisateurs\n Pour modifier le nombres de polarités, d'utilisateurs ou de qbits, modifier les variables nbQ, pola et nbUsers dans le code com.go")
	}

}
