package models

import (
	"github.com/google/uuid"
)

// ClientOrder represents a clients request for a trade order.
type ClientOrder struct {
	Id        uuid.UUID
	Symbol    string
	OrderSize string
	Price     string
}
