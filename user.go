package main

type User struct {
	qbits []qbit
}

func (u User) getQbits() []qbit {
	return u.qbits
}

func (u User) setQbits(qbits []qbit) {
	u.qbits = qbits
}
