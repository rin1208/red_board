package domain

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Shelter struct {
	gorm.Model
	Name    string `json:"name"`
	Comment string `json:"comment"`
	Time    string `json:"time"`
}
