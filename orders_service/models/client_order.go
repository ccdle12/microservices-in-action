package models

import (
	"github.com/jinzhu/gorm"
)

type ClientOrder struct {
	gorm.Model
	Id        string
	Symbol    string
	OrderSize string
	Price     string
}
