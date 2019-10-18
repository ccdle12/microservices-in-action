package models

import (
	"testing"
)

// Test that a client order can be created.
func TestClientOrder(t *testing.T) {
	// TODO(ccdle12):
	// UserId should be a UUID
	// Symbol an Enum?
	// Amount a string? or renamed to order_size?
	_ = &ClientOrder{UserId: "123", Symbol: "BTC", Amount: "1234"}
}
