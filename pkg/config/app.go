package config

import (
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "root/bookrest?charset=utf8&Time=true&loc=Local")

	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
