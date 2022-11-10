package main

import (
	"math/rand"
	"time"
)

func test75(nbQ int, pola int) int {
	alice := User{make([]qbit, nbQ)}
	bob := User{make([]qbit, nbQ)}
	cpt := 0
	//generate random bits
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < nbQ; i++ {
		r := rand.Intn(pola)
		value := rand.Intn(2)
		alice.getQbits()[i] = qbit{value, r}
		basis := rand.Intn(pola)
		bob.getQbits()[i] = qbit{alice.getQbits()[i].measure(basis, pola), alice.getQbits()[i].getPolarisation()}
		if bob.getQbits()[i].getValue() == alice.getQbits()[i].getValue() {
			cpt++
		}
	}
	return (cpt * 100) / nbQ
}
func test62(nbQ int, pola int) int {
	alice := User{make([]qbit, nbQ)}
	bob := User{make([]qbit, nbQ)}
	eve := User{make([]qbit, nbQ)}
	cpt := 0
	//generate random bits
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < nbQ; i++ {
		r := rand.Intn(pola)
		value := rand.Intn(2)
		alice.getQbits()[i] = qbit{value, r}
		basisEve := rand.Intn(pola)
		basisBob := rand.Intn(pola)
		eve.getQbits()[i] = qbit{alice.getQbits()[i].measure(basisEve, pola), alice.getQbits()[i].getPolarisation()}
		bob.getQbits()[i] = qbit{eve.getQbits()[i].measure(basisBob, pola), eve.getQbits()[i].getPolarisation()}
		if bob.getQbits()[i].getValue() == alice.getQbits()[i].getValue() {
			cpt++
		}
	}
	return (100 * cpt) / nbQ
}

func test3users(nbQ int, pola int) int {
	alice := User{make([]qbit, nbQ)}
	bob := User{make([]qbit, nbQ)}
	eve := User{make([]qbit, nbQ)}
	someone := User{make([]qbit, nbQ)}
	cpt := 0
	//generate random bits
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < nbQ; i++ {
		r := rand.Intn(pola)
		value := rand.Intn(2)
		alice.getQbits()[i] = qbit{value, r}
		basisEve := rand.Intn(pola)
		basisBob := rand.Intn(pola)
		basisS := rand.Intn(pola)
		eve.getQbits()[i] = qbit{alice.getQbits()[i].measure(basisEve, pola), alice.getQbits()[i].getPolarisation()}
		someone.getQbits()[i] = qbit{eve.getQbits()[i].measure(basisS, pola), eve.getQbits()[i].getPolarisation()}
		bob.getQbits()[i] = qbit{someone.getQbits()[i].measure(basisBob, pola), someone.getQbits()[i].getPolarisation()}
		if bob.getQbits()[i].getValue() == alice.getQbits()[i].getValue() {
			cpt++
		}
	}
	return (100 * cpt) / nbQ
}
