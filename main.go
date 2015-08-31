package main

import (
	"github.com/iaas-engine/interfaces"
	"github.com/iaas-engine/usecases"	
	"github.com/gorilla/mux"
	"github.com/codegangsta/negroni"
)


func main(){

	
	handler := interfaces.WebServiceHandler {
		EngineInteractor: usecases.EngineInteractor{},
	}
	
	r := mux.NewRouter()
	subrouter := r.PathPrefix("/").Subrouter()
	subrouter.HandleFunc("/create", handler.CreateConfig).Methods("POST")

	n := negroni.Classic()
	n.UseHandler(r)

	n.Run(":7002")




}
