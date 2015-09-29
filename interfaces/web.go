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
	CreateRepo(server domain.Server, user, token string) error
}

type WebServiceHandler struct {
	EngineInteractor EngineInteractor
}

type User struct {
	ID          int
	Username    string
	AccessToken string
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

	server := getConfig(req)
	
	handler.EngineInteractor.CreateZip(server, res)

}

func (handler WebServiceHandler) CreateRepo(res http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	
	token := req.Header.Get("gitToken")
	user := req.Header.Get("gitUser")
	
	res.Header().Set("Content-Type", "application/json")
	server := getConfig(req)
	
	err := handler.EngineInteractor.CreateRepo(server, user, token)
	if err != nil{
		res.WriteHeader(http.StatusInternalServerError)
	} else {
		res.WriteHeader(http.StatusOK)
	}

}
