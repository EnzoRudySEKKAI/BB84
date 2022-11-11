// qbit bb84 protocol
package main

import (
	"math/rand"
	"os"
	"time"
)

func testNusers(nbQ int, pola int, n int) int {
	listUsers := make([]User, n)
	alice := User{make([]qbit, nbQ)}
	listUsers[0] = alice

	if n > 2 {
		for i := 1; i < n-1; i++ {
			listUsers[i] = User{make([]qbit, nbQ)}
		}
	}

	bob := User{make([]qbit, nbQ)}
	listUsers[n-1] = bob

	cpt := 0
	//generate random bits
	rand.Seed(time.Now().UTC().UnixNano())

	for j := 0; j < n-1; j++ {
		for i := 0; i < nbQ; i++ {
			r := rand.Intn(pola)

			//remplissage alice
			if j == 0 {
				r = rand.Intn(pola)
				listUsers[j].getQbits()[i].setValue(r)
			}

			r = rand.Intn(pola)
			basisTemp := rand.Intn(pola)

			listUsers[j+1].getQbits()[i] = qbit{listUsers[j].getQbits()[i].measure(basisTemp), listUsers[j].getQbits()[i].getPolarisation()}

		}

	}
	for i := 0; i < nbQ; i++ {
		if bob.getQbits()[i].getValue() == alice.getQbits()[i].getValue() {
			cpt++
		}
	}
	return (100 * cpt) / nbQ
}

func main() {
	if os.Args[1] == "server" {
		Server()
	} else if os.Args[1] == "client" {
		Client()
	}
}
