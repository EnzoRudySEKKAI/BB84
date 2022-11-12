package main

import "math/rand"

type Qbit struct {
	value        int
	polarisation int
}

func (q Qbit) getValue() int {
	return q.value
}

func (q Qbit) getPolarisation() int {
	return q.polarisation
}

func (q Qbit) setPolarisation(polarisation int) {
	q.polarisation = polarisation
}

func (q Qbit) setValue(value int) {
	q.value = value
}

func (q Qbit) measure(basis int) int {
	if basis == q.getPolarisation() {
		return q.getValue()
	} else {
		return rand.Intn(2)
	}
}
