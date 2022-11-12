// Qbit bb84 protocol
package main

import (
	"fmt"
	"os"
)

func main() {
	//number of qbits
	nbQ := 100
	//number of polarities
	pola := 2
	//number of users
	nbUsers := 3
	//number of tests
	nTest := 100000
	// % of success
	var cptfi float64 = 0

	if len(os.Args) > 1 {
		if os.Args[1] == "server" {
			Server()

		} else if os.Args[1] == "client" {
			Client()

		} else if os.Args[1] == "75" {
			fmt.Println("Test with:", pola, "polarities and ", nbQ, " qbits. (", nTest, " tests)")
			for n := 0; n < nTest; n++ {
				cptfi += float64(test75(nbQ, pola))
			}
			var cptfinal = cptfi / float64(nTest)
			fmt.Println("2 users without eve :", cptfinal)

		} else if os.Args[1] == "62" {
			for n := 0; n < nTest; n++ {
				cptfi += float64(test62(nbQ, pola))
			}
			cptfi = cptfi / float64(nTest)
			fmt.Println("2 users with eve :", cptfi)

		} else if os.Args[1] == "4users" {
			for n := 0; n < nTest; n++ {
				cptfi += float64(test4users(nbQ, pola))
			}
			cptfi = cptfi / float64(nTest)
			fmt.Println("2 users with eve and someone else:", cptfi)

		} else if os.Args[1] == "Nusers" {
			for n := 0; n < nTest; n++ {
				cptfi += float64(testNusers(nbQ, pola, nbUsers))
			}
			cptfi = cptfi / float64(nTest)
			fmt.Println(nbUsers, " users and ", pola, "polarities: ", cptfi)
		}
	} else {
		fmt.Println("argument list:\n server : launch the serveur\n client : launch the client\n 75 : launch the 75% test\n 62 : launch the 62% test\n 4users : launch the test with 4 users\n Nusers : launch the test with N users\n You can change the number of Users, Polarities, Qbits and Tests with the variables nbUser, pola, nbQ and nTest in com.go")
	}

}
