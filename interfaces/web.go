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
	CreateConf(server domain.Server, zipFile io.Writer)
}

type WebServiceHandler struct {
	EngineInteractor EngineInteractor
}

func (handler WebServiceHandler) CreateConfig(res http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	decoder := json.NewDecoder(req.Body)
	mainJSON := domain.Myjson{}
	if e := decoder.Decode(&mainJSON); e != nil {
		log.Fatal(e)
	}

	server := domain.Server{}
	server = mainJSON.Server

	res.Header().Set("Content-Type", "application/zip")
	res.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", "puppet.zip"))
	handler.EngineInteractor.CreateConf(server, res)

}
