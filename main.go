package main

import (
	"fmt"
	"net/http"

	"../project/lib"
)

func init() {
	lib.DBConnection()
}

func main() {
	route := lib.Router()
	fmt.Println("Http listning at port:", lib.Service_Config.Port)
	http.ListenAndServe(":"+lib.Service_Config.Port, route)
}
