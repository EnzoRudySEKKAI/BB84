package main

import "math/rand"

type qbit struct {
	value        int
	polarisation int
}

func (q qbit) getValue() int {
	return q.value
}

func (q qbit) getPolarisation() int {
	return q.polarisation
}

func (q qbit) setPolarisation(polarisation int) {
	q.polarisation = polarisation
}

func (q qbit) setValue(value int) {
	q.value = value
}

func (q qbit) measure(basis int) int {
	if basis == q.getPolarisation() {
		return q.getValue()
	} else {
		return rand.Intn(2)
	}
}
