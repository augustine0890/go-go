```go
package main

import (
	"fmt"

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

	for _, user := range users {
		db.Create(&user)
	}
	u := User{Username: "tmacmillan"}
	// Querying
	// db.First(&u)
	// db.Last(&u)

	// Updating
	db.Where(&u).First(&u)
	u.LastName = "Beeblebrox"
	db.Save(&u)

	user := User{}
	db.Where(&u).First(&user)

	// Deleting
	db.Where(&User{Username: "mrobot"}).Delete(&User{})

	fmt.Println(user)
}

type User struct {
	ID        int
	Username  string
	FirstName string
	LastName  string
}

var users []User = []User{
	{Username: "august", FirstName: "Augustine", LastName: "Nguyen"},
	{Username: "fperfact", FirstName: "Ford", LastName: "Prefect"},
	{Username: "tmacmillan", FirstName: "Tricia", LastName: "MacMillan"},
	{Username: "mrobot", FirstName: "Marvin", LastName: "Robot"},
}
```