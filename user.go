package main

type User struct {
	qbits []Qbit
}

func (u User) getQbits() []Qbit {
	return u.qbits
}

func (u User) setQbits(qbits []Qbit) {
	u.qbits = qbits
}
