package entity

type User interface {
	ID() int
	FirstName() string
	LastName() string
	CarTag() string
	Phone() string
	Password() string
}
