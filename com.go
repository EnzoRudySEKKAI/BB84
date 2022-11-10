// qbit bb84 protocol
package main

import (
	"fmt"
	"math/rand"
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

			listUsers[j+1].getQbits()[i] = qbit{listUsers[j].getQbits()[i].measure(basisTemp, pola), listUsers[j].getQbits()[i].getPolarisation()}

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
	var cptfi float64 = 0
	nTest := 20

	pola := 2
	nQ := 1000
	nb_personnes := 2

	fmt.Println("Test avec:", pola, "polarités et ", nQ, " qbits. (", nTest, " tests)")
	for n := 0; n < nTest; n++ {
		cptfi += float64(test75(nQ, pola))
	}
	var cptfinal = cptfi / float64(nTest)
	fmt.Println("2 personnes sans eve :", cptfinal)

	cptfi = 0
	for n := 0; n < nTest; n++ {
		cptfi += float64(test62(nQ, pola))
	}
	cptfinal = cptfi / float64(nTest)
	fmt.Println("2 personnes avec eve :", cptfinal)

	cptfi = 0
	for n := 0; n < nTest; n++ {
		cptfi += float64(test3users(nQ, pola))
	}
	cptfinal = cptfi / float64(nTest)
	fmt.Println("2 personnes avec eve et quelqu'un d'autre:", cptfinal)

	cptfi = 0
	for n := 0; n < nTest; n++ {
		cptfi += float64(testNusers(nQ, pola, nb_personnes))
	}
	cptfinal = cptfi / float64(nTest)
	fmt.Println("N personnes:", cptfinal)
}
