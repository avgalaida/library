package domain

import (
	"github.com/avgalaida/library/infrastructure/event_sourcing"
)

type Book struct {
	Base    event_sourcing.BasedAggregate
	Status  string
	Title   string
	Authors string
}
