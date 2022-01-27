package lib

import "github.com/jinzhu/gorm"

type Postgres struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Db_Name  string `yaml:"db_name"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
}

type Service struct {
	Port string `yaml:"port"`
}

type patient struct {
	gorm.Model
	Name     string
	Age      int
	Email    string
	Contact  int64
	Password string
}
