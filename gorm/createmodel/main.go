package main

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	ID       uint      `gorm:"primary_key"`
	username string    `gorm:"type:varchar(255)"`
	createAt time.Time `json:"create_at"`
	updateAt time.Time `json:"update_at"`
	postno   string    `gorm:"type:varchar(255)"`
}

type Post struct {
	ID        uint      `gorm:"primary_key"`
	content   string    `gorm:"type:text"`
	userid    uint      `gorm:"type:varchar(255)"`
	createAt  time.Time `json:"create_at"`
	updateAt  time.Time `json:"update_at"`
	commentno string    `gorm:"type:varchar(255)"`
}

type Comments struct {
	ID       uint      `gorm:"primary_key"`
	content  string    `gorm:"type:text"`
	postid   uint      `gorm:"type:varchar(255)"`
	userid   uint      `gorm:"type:varchar(255)"`
	username string    `gorm:"type:varchar(255)"`
	createAt time.Time `json:"create_at"`
	updateAt time.Time `json:"update_at"`
}

func main() {
	dsn := "gorm:gorm1234!@tcp(127.0.0.1:3306)/gorm?charset=utf8"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Fatalf("init db failed", err)
	}

	err = db.AutoMigrate(&User{}, &Post{}, &Comments{})
	if err != nil {
		log.Fatalf("init db failed", err)
	}
	log.Printf("create table success")
}
