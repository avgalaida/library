package domain

type Book struct {
	Base    BasedAggregate
	Status  string
	Title   string
	Authors string
}
