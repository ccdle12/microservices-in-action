package models

import (
	"github.com/jinzhu/gorm"
)

type ClientOrder struct {
	gorm.Model
	UserId string
	Symbol string
	Amount string
}
