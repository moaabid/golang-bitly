package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

type Bitly struct {
	ID       uint64 `json:"id" gorm:"primaryKey"`
	Redirect string `json:"redirect" gorm:"unique;not null"`
	Bitly     string `json:"bitly" gorm:"unique;not null"`
	Clicked  uint64 `json:"clicked"`
	Random   bool   `json:"random"`
}

func ConnectDB() {

	var err error
	dsn := "host=localhost user=postgres password=password dbname=bitly port=5432 sslmode=disable"

	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	err = Db.AutoMigrate(&Bitly{})

	if err != nil {
		fmt.Println(err.Error())
	}
}
