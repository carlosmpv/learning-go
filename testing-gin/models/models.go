package models

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Name  string  `form:"name"`
	Price float64 `form:"price"`
}
