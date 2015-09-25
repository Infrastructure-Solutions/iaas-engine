package main

import (
	"flag"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/iaas-engine/interfaces"
	"github.com/iaas-engine/usecases"
	"github.com/iaas-engine/infraestructure"
)

var defaultPath = "/etc/iaas-engine.conf"
var confFilePath = flag.String("conf", defaultPath, "Custom Path for configuration file")

func main() {

	flag.Parse()

	config, err := infraestructure.GetConfiguration(*confFilePath)
	if err != nil {
		panic("Cannot parse configuration")
	}

	writer, err := infraestructure.NewFileWriter(config.TemplatesPath)
	if err != nil {
		panic("Cannot set templates path")
	}

	interactor, err := usecases.NewEngineInteractor(config.FilesPath, writer)
	if err != nil {
		panic("Cannot set files path")
	}
	
	handler := interfaces.WebServiceHandler{
		EngineInteractor: interactor,
	}

	r := mux.NewRouter()
	subrouter := r.PathPrefix("/iaas").Subrouter()
	subrouter.HandleFunc("/create", handler.CreateConfig).Methods("POST")

	n := negroni.Classic()
	n.UseHandler(r)

	n.Run(":7002")

}
