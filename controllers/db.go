package controllers

import (

  _ "github.com/go-sql-driver/mysql"
  //v "github.com/spf13/viper"
  "github.com/jinzhu/gorm"
  "log"

  "../models"

)

type DBController struct {
	DB *gorm.DB
}

func (dc *DBController) InitDB() {
	var err error

	dc.DB, err = gorm.Open("mysql","root:12345@/todolist?charset=utf8&parseTime=True")
	if err != nil {
		log.Fatalf("Error when connect database, the error is '%v'", err)
	}
	dc.DB.LogMode(true)
}

func (dc *DBController) GetDB() *gorm.DB {
  return dc.DB
}

func (dc *DBController) InitSchema(){
  
  dc.DB.AutoMigrate(&models.Users{})
  dc.DB.AutoMigrate(&models.TodoList{})
}


