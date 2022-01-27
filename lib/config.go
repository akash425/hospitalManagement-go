package lib

import (
	"fmt"
	"io/ioutil"

	"github.com/go-yaml/yaml"
)

var Postgres_Config Postgres
var Service_Config Service

func init() {
	err := GetConfiguration()
	if err != nil {
		fmt.Println(err)
	}
}

func GetConfiguration() error {

	yamlFile, err := ioutil.ReadFile("config/postgres.yaml")
	if err != nil {
		fmt.Println("Error! ReadFile[GetPostgresConfigration]:", err)
		return err
	}

	err = yaml.Unmarshal(yamlFile, &Postgres_Config)
	if err != nil {
		fmt.Println("Error! Unmarshal[GetPostgresConfigration]:", err)
		return err
	}

	yamlFile, err = ioutil.ReadFile("config/service.yaml")
	if err != nil {
		fmt.Println("Error! ReadFile[GetPostgresConfigration]:", err)
		return err
	}

	err = yaml.Unmarshal(yamlFile, &Service_Config)
	if err != nil {
		fmt.Println("Error! Unmarshal[GetPostgresConfigration]:", err)
		return err
	}
	return err
}
