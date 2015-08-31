package usecases

import (
	"github.com/iaas-engine/infraestructure"
	"github.com/iaas-engine/domain"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"archive/zip"
	"time"
	"strings"
	"path"
)

type EngineInteractor struct {
}

var configs = []struct{
	Name, Path string
}{
	{"provisioner.sh", "provisioner.sh"},
	{"Vagrantfile", "Vagrantfile"},
	{"hiera.yaml","hiera.yaml"},
	{".gitmodules",".gitmodules"},
	{"environment.conf","environments/tequilaware/environment.conf"},
	{"hosts.yaml","hieradata/hosts.yaml"},
	{"site.pp","manifests/site.pp"},
}

func (interactor EngineInteractor) CreateConf(server domain.Server, zipFile io.Writer) {
	var hieraClasses = []string{}
	var Files = []domain.File{}
	
	packages := []domain.Package{}
	packages = server.GetPackages()
	className := server.GetHostname()
	
	hieraClasses = append(hieraClasses, className)
	
	content := createPackages(packages, hieraClasses)
	
	manifest := domain.Manifest{ClassName:className, Content:content}
	
	path := "environments/tequilaware/modules/web/manifests/init.pp"
	doc, error := infraestructure.WriteClass(manifest)
	if error != nil {
		fmt.Println(error)
	}

	file := domain.File{path,doc}
	Files = append(Files, file)
	
	path = "hieradata/tequilaware/node/web.yaml"
	doc1, error := infraestructure.WriteHiera(hieraClasses)
	if error != nil {
		fmt.Println(error)
	}
	file1 := domain.File{path,doc1}
	Files = append(Files, file1)
	
	Files = append(Files, getPuppetFiles()...)
	
	writeZip(zipFile, Files)

}

func getPuppetFiles() []domain.File {
	paths := "infraestructure/files/puppet"
	var files = []domain.File{}
	for _, file := range configs {
		content, err := infraestructure.ReadFile(path.Join(paths,file.Name))
		if err != nil {
			log.Print(err)
		}
		files = append(files, domain.File{file.Path, content})
	}
	return files
}

func createPackages(packages []domain.Package, hieraClasses []string) string { 
	var manifestContent string
	for _, elem := range packages {
		if elem.Config != nil {
			switch {
			case elem.Name == "nginx":
				hieraClasses = append(hieraClasses, elem.Name)
				nginxConf := domain.NginxConfig{}
				json.Unmarshal(elem.Config, &nginxConf)
				doc, error := infraestructure.WriteConf(nginxConf, elem.Name)
				if error != nil {
					fmt.Println(error)
				}
				manifestContent += doc
			default : fmt.Println("Uknown config")
			}
		} else {
			doc, error := infraestructure.WritePackages(elem)
			if error != nil {
				fmt.Println(error)
			}
			manifestContent += doc
		}
	}
	return manifestContent
}

func writeZip(zipFile io.Writer, Files []domain.File) {
	w := zip.NewWriter(zipFile)
	for _, file := range Files {
		header := &zip.FileHeader{
			Name:         file.Path,
			Method:       zip.Store,
			ModifiedTime: uint16(time.Now().UnixNano()),
			ModifiedDate: uint16(time.Now().UnixNano()),
		}
		fw, err := w.CreateHeader(header)
		if err != nil {
			log.Fatal(err)
		}
		reader := strings.NewReader(file.Content)
		if _, err = io.Copy(fw, reader); err != nil {
			log.Fatal(err)
		}
	}
	err := w.Close()
	if err != nil {
		log.Fatal(err)
	}
}
