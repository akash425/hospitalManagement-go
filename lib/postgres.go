package lib

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func DBConnection() *gorm.DB {
	DSN := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", Postgres_Config.User, Postgres_Config.Password, Postgres_Config.Db_Name)
	db, err := gorm.Open("postgres", DSN)
	if err != nil {
		fmt.Println("Error", err)
	}
	db.AutoMigrate(&patient{})
	return db
}

func find(person patient) string {
	var result patient
	db := DBConnection()
	db.Where("Email=?", person.Email).Find(&result)
	defer db.Close()
	return result.Email
}

func create(person patient) {
	db := DBConnection()
	db.Create(&person)
	defer db.Close()
}

func delete(person patient) {
	var result patient
	db := DBConnection()
	db.Where("Email=?", person.Email).Delete(&result)
	defer db.Close()
}

func update(person patient) {
	db := DBConnection()
	db.Model(&person).Where("Email=?", person.Email).Update("Age", person.Age)
	defer db.Close()
}

func show(person []patient) []patient {
	db := DBConnection()
	db.Find(&person)
	return person
}
func login(person patient) patient {
	var result patient
	db := DBConnection()
	db.Where("Email=?", person.Email).Find(&result)
	defer db.Close()
	return result
}
