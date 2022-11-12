package main

import (
	"math/rand"
	"time"
)

func test75(nbQ int, pola int) int {
	//create users
	alice := User{make([]Qbit, nbQ)}
	bob := User{make([]Qbit, nbQ)}
	cpt := 0
	//generate random bits
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < nbQ; i++ {
		r := rand.Intn(pola)
		value := rand.Intn(2)
		alice.getQbits()[i] = Qbit{value, r}
		basis := rand.Intn(pola)
		bob.getQbits()[i] = Qbit{alice.getQbits()[i].measure(basis), alice.getQbits()[i].getPolarisation()}
		if bob.getQbits()[i].getValue() == alice.getQbits()[i].getValue() {
			cpt++
		}
	}
	return (cpt * 100) / nbQ
}
func test62(nbQ int, pola int) int {
	alice := User{make([]Qbit, nbQ)}
	bob := User{make([]Qbit, nbQ)}
	eve := User{make([]Qbit, nbQ)}
	cpt := 0
	//generate random bits
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < nbQ; i++ {
		r := rand.Intn(pola)
		value := rand.Intn(2)
		alice.getQbits()[i] = Qbit{value, r}
		basisEve := rand.Intn(pola)
		basisBob := rand.Intn(pola)
		eve.getQbits()[i] = Qbit{alice.getQbits()[i].measure(basisEve), alice.getQbits()[i].getPolarisation()}
		bob.getQbits()[i] = Qbit{eve.getQbits()[i].measure(basisBob), eve.getQbits()[i].getPolarisation()}
		if bob.getQbits()[i].getValue() == alice.getQbits()[i].getValue() {
			cpt++
		}
	}
	return (100 * cpt) / nbQ
}

func test4users(nbQ int, pola int) int {
	alice := User{make([]Qbit, nbQ)}
	bob := User{make([]Qbit, nbQ)}
	eve := User{make([]Qbit, nbQ)}
	someone := User{make([]Qbit, nbQ)}
	cpt := 0
	//generate random bits
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < nbQ; i++ {
		r := rand.Intn(pola)
		value := rand.Intn(2)
		alice.getQbits()[i] = Qbit{value, r}
		basisEve := rand.Intn(pola)
		basisBob := rand.Intn(pola)
		basisS := rand.Intn(pola)
		eve.getQbits()[i] = Qbit{alice.getQbits()[i].measure(basisEve), alice.getQbits()[i].getPolarisation()}
		someone.getQbits()[i] = Qbit{eve.getQbits()[i].measure(basisS), eve.getQbits()[i].getPolarisation()}
		someone.getQbits()[i] = Qbit{eve.getQbits()[i].measure(basisS), eve.getQbits()[i].getPolarisation()}
		bob.getQbits()[i] = Qbit{someone.getQbits()[i].measure(basisBob), someone.getQbits()[i].getPolarisation()}
		if bob.getQbits()[i].getValue() == alice.getQbits()[i].getValue() {
			cpt++
		}
	}
	return (100 * cpt) / nbQ
}

func testNusers(nbQ int, pola int, n int) int {
	listUsers := make([]User, n)
	alice := User{make([]Qbit, nbQ)}
	listUsers[0] = alice

	if n > 2 {
		for i := 1; i < n-1; i++ {
			listUsers[i] = User{make([]Qbit, nbQ)}
		}
	}

	bob := User{make([]Qbit, nbQ)}
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

			listUsers[j+1].getQbits()[i] = Qbit{listUsers[j].getQbits()[i].measure(basisTemp), listUsers[j].getQbits()[i].getPolarisation()}

		}

	}
	for i := 0; i < nbQ; i++ {
		if bob.getQbits()[i].getValue() == alice.getQbits()[i].getValue() {
			cpt++
		}
	}
	return (100 * cpt) / nbQ
}
