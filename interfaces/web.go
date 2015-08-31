package interfaces

import (
	"github.com/iaas-engine/domain"	
	"net/http"
	"encoding/json"
	"log"
	"io"
	"fmt"
)

type EngineInteractor interface {
	CreateConf(server domain.Server, zipFile io.Writer)
}

type WebServiceHandler struct {
	EngineInteractor EngineInteractor
}

func (handler WebServiceHandler) CreateConfig(res http.ResponseWriter, req *http.Request){
	defer req.Body.Close()

	decoder := json.NewDecoder(req.Body)
	server := domain.Myjson{}
	if e := decoder.Decode(&server); e != nil {
		log.Fatal(e)
	}

	servidor := domain.Server{}
	servidor = server.GetServer()

	res.Header().Set("Content-Type", "application/zip")
	res.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", "algo"))
	handler.EngineInteractor.CreateConf(servidor, res)
	
}
