package models

import (
	"github.com/google/uuid"
	"testing"
)

// Test that a client order can be created.
func TestClientOrder(t *testing.T) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		t.Errorf("Failed to generate uuid: %v", err)
	}

	_ = &ClientOrder{
		Id:        uuid,
		Symbol:    "BTC",
		OrderSize: "1234",
		Price:     "123",
	}
}
