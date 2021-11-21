package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=database port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	db.Migrator().DropTable(&User{})
	db.AutoMigrate(&User{})
}

type User struct {
	gorm.Model
	Username  string `gorm:"type:varchar(15);unique_index"`
	FirstName string `gorm:"size:100"`
	LastName  string
}

// func (u User) TableName() string {
// return "stakeholders"
// }
