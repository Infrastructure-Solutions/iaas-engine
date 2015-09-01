package main

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/iaas-engine/interfaces"
	"github.com/iaas-engine/usecases"
)

func main() {

	handler := interfaces.WebServiceHandler{
		EngineInteractor: usecases.EngineInteractor{},
	}

	r := mux.NewRouter()
	subrouter := r.PathPrefix("/iaas").Subrouter()
	subrouter.HandleFunc("/create", handler.CreateConfig).Methods("POST")

	n := negroni.Classic()
	n.UseHandler(r)

	n.Run(":7002")

}
