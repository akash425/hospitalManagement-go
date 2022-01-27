package lib

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../helper"
	"github.com/dgrijalva/jwt-go"
)

var per string
var tokenstring string

//log_status := false

type Claims struct {
	Email string
	jwt.StandardClaims
}

func Create_patient(w http.ResponseWriter, r *http.Request) {
	var person patient
	var result patient
	json.NewDecoder(r.Body).Decode(&person)
	result.Email = find(person)
	if result.Email == "" {
		/*pswd := result.Password
		H_Password, err := helper.HashPassword(pswd)
		if err != nil {
			log.Println(err)
		}
		per = string(H_Password)
		person.Password = per
		*/
		create(person)
		helper.WriteResponse(w, "Registerd", http.StatusOK)
	} else {
		helper.WriteResponse(w, "User already exists", http.StatusConflict)
	}
}

func Delete_patient(w http.ResponseWriter, r *http.Request) {
	var person patient
	var result patient
	json.NewDecoder(r.Body).Decode(&person)
	fmt.Println(person.Email)
	result.Email = find(person)
	if result.Email == "" {
		helper.WriteResponse(w, "No user with this email", http.StatusBadRequest)
	} else {
		delete(person)
		helper.WriteResponse(w, "Deleted", http.StatusOK)
	}
}

func Update_patient(w http.ResponseWriter, r *http.Request) {
	var person patient
	var result patient
	json.NewDecoder(r.Body).Decode(&person)
	fmt.Println(person.Email)
	result.Email = find(person)
	if result.Email == "" {
		helper.WriteResponse(w, "No user with this email", http.StatusBadRequest)
	} else {
		update(person)
		helper.WriteResponse(w, "updated", http.StatusOK)
	}
}

func Show_patient(w http.ResponseWriter, r *http.Request) {
	var person []patient
	getToken := r.Header.Get("Authorization")
	fmt.Println(getToken)
	res := helper.Validation(getToken)
	if res == true {
		println("ok")
		person = show(person)
	} else {
		helper.WriteResponse(w, "Validation Failed", http.StatusOK)
	}
}

func Login_patient(w http.ResponseWriter, r *http.Request) {
	var person patient
	var result patient
	json.NewDecoder(r.Body).Decode(&person)
	result = login(person)
	if result.Email == "" {
		helper.WriteResponse(w, "No user with this email", http.StatusBadRequest)
	} else {
		//match := helper.CheckPasswordHash(person.Password, result.Password)
		match := person.Password == result.Password
		fmt.Println(match)
		if match == true {
			tokenstring = helper.Loginvalid(result.Email)
			//log_status = true
			fmt.Println(tokenstring)
			helper.WriteResponse(w, "Logged In", http.StatusOK)
		} else {
			helper.WriteResponse(w, "Wrong Password", http.StatusBadRequest)
		}
	}
}
