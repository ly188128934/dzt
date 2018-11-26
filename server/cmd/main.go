package main

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

var db * gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql". "user:password/database?charset=utf-8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
}

type clue struct {
	location	string	`gorm:"type:varchar(10);not null;"`
	clue		string	`gorm:"type:varchar(256);not null;"`
	status		bool
}


func main() {
	fmt.Println("1")
}
