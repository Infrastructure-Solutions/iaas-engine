package interfaces

import (
	"encoding/json"
	"fmt"
	"github.com/iaas-engine/domain"
	"io"
	"log"
	"net/http"
)

type EngineInteractor interface {
	CreateZip(server domain.Server, zipFile io.Writer)
	CreateRepo(server domain.Server, files io.Writer)
}

type WebServiceHandler struct {
	EngineInteractor EngineInteractor
}

func getConfig(req *http.Request) domain.Server {
	decoder := json.NewDecoder(req.Body)
	mainJSON := domain.Myjson{}
	if e := decoder.Decode(&mainJSON); e != nil {
		log.Fatal(e)
	}

	server := domain.Server{}
	server = mainJSON.Server

	return server
}

func (handler WebServiceHandler) CreateZip(res http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	
	res.Header().Set("Content-Type", "application/zip")
	res.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", "puppet.zip"))
	handler.EngineInteractor.CreateConf(server, res)
	server := getConfig(req)
	
	handler.EngineInteractor.CreateZip(server, res)

}

func (handler WebServiceHandler) CreateRepo(res http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	
	res.Header().Set("Content-Type", "application/json")
	server := getConfig(req)
	
	handler.EngineInteractor.CreateRepo(server, res)

}
