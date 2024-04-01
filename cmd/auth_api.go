package main

import (
	"fmt"
	"log"

	"github.com/nguyentrunghieu15/auth-api-vcs-pj/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func main() {

	dsn := "host=localhost port=5432 dbname=on_demand_services_db user=hiro password=1 connect_timeout=10 sslmode=prefer"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Can't connect DB")
	}
	fmt.Print(db)
	var user model.User
	var avatar model.Avatar

	// works because destination struct is passed in
	db.Preload(clause.Associations).First(&user)
	// SELECT * FROM `users` ORDER BY `users`.`id` LIMIT 1

	fmt.Println(user)

	db.First(&avatar)
	fmt.Println(avatar)
}
