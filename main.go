package main

import (
	"net/http"

	"github.com/Diegovalen47/mutants/db"
	"github.com/Diegovalen47/mutants/models"
	"github.com/Diegovalen47/mutants/routes"
	"github.com/gorilla/mux"
)

func main() {

	db.DBConnection()

	db.DB.AutoMigrate(&models.Dna{})

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)
	r.HandleFunc("/mutant", routes.PostDnaHandler).Methods("POST")
	r.HandleFunc("/stats", routes.GetDnaStatsHandler).Methods("GET")

	http.ListenAndServe(":8080", r)
}
