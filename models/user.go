package models

import (
	"github.com/jinzhu/gorm"
)

//
type User struct {
	gorm.Model
	Name     string  `json:"name"`
	HigherID uint    `json:"higher_id"`
	Higher   User    `json:"higher"`
	Agent    bool    `json:"agent"`
	Money    float64 `json:"money"`
}
