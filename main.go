package main

import (
	"flag"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/iaas-engine/interfaces"
	"github.com/iaas-engine/usecases"
	"github.com/iaas-engine/infraestructure"
)

var defaultPath = "iaas-engine.conf"
var confFilePath = flag.String("conf", defaultPath, "Custom Path for configuration file")

func main() {

	flag.Parse()

	config, err := infraestructure.GetConfiguration(*confFilePath)
	if err != nil {
		panic("Cannot parse configuration")
	}

	writer, err := infraestructure.NewFileWriter(config.TemplatesPath, config.FilesPath)
	if err != nil {
		panic("Cannot set templates path")
	}

	git := interfaces.NewGit("http://localhost:7000")
	
	interactor, err := usecases.NewEngineInteractor(writer, git)
	if err != nil {
		panic("Cannot set files path")
	}
	
	handler := interfaces.WebServiceHandler{
		EngineInteractor: interactor,
	}

	r := mux.NewRouter()
	subrouter := r.PathPrefix("/iaas").Subrouter()
	createRouter := subrouter.PathPrefix("/create").Subrouter()
	createRouter.HandleFunc("/zip", handler.CreateZip).Methods("POST")
	createRouter.HandleFunc("/repo", handler.CreateRepo).Methods("POST")

	n := negroni.Classic()
	n.UseHandler(r)

	n.Run(":7002")

}
