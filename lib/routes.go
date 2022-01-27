package lib

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Router() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/create_patient", Create_patient).Methods("POST")
	r.HandleFunc("/delete_patient", Delete_patient).Methods("DELETE")
	r.HandleFunc("/update_patient", Update_patient).Methods("PUT")
	r.HandleFunc("/show_patient", Show_patient).Methods("GET")
	r.HandleFunc("/login_patient", Login_patient).Methods("POST")
	http.Handle("/", r)
	return r
}
