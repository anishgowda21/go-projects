package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


var (
	db *gorm.DB
)

func Connect(){
	d, err := gorm.Open("mysql","root:mysql@tcp(127.0.0.1:3306)/go_bookstore?charset=utf8mb4&parseTime=True&loc=Local")

	if err!=nil{
		panic(err)
	}else{
		fmt.Println("Connected to db")
	}

	db = d
}


func GetDB() *gorm.DB{
	return db
}
