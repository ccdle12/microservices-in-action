package models

import (
	"testing"
)

// Test that a client order can be created.
func TestClientOrder(t *testing.T) {
	_ = &ClientOrder{
		Id:        "123",
		Symbol:    "BTC",
		OrderSize: "1234",
		Price:     "123",
	}
}
