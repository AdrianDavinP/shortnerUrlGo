package main

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
	"mraladin/shortnerUrl/src/Process"
	"net/http"
)

var db *gorm.DB
var err error

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/{id}", Process.HitUrlPendek).Methods("GET")
	router.HandleFunc("/auth/getToken", Process.GetToken).Methods("GET")

	s := router.PathPrefix("/auth").Subrouter()
	s.Use(Process.JwtVerify)
	s.HandleFunc("/daftarUrl", Process.DaftarUrl).Methods("PUT")
	s.HandleFunc("/ubahUrl", Process.UbahUrl).Methods("PUT")
	s.HandleFunc("/delete", Process.Delete).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":23456", router))
}
