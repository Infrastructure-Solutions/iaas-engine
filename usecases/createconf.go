package usecases

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"github.com/iaas-engine/domain"
	"io"
	"io/ioutil"
	"log"
	"path"
	"strings"
	"time"
)

type EngineInteractor struct {
	filesPath  string
	FileWriter FileWriter
}

func NewEngineInteractor(filesPath string, fileWriter FileWriter) (*EngineInteractor, error) {
	interactor := &EngineInteractor{
		filesPath: filesPath,
		FileWriter: fileWriter,
	}
	return interactor, nil
}

type FileWriter interface {
	WriteTemplate(conf interface{}, pack string) (string, error)
}

type Template struct {
	Name   string
	Config interface{}
	Path   string
}

var configs = []struct {
	Name, Path string
}{
	{"provisioner.sh", "provisioner.sh"},
	{"Vagrantfile", "Vagrantfile"},
	{"hiera.yaml", "hiera.yaml"},
	{".gitmodules", ".gitmodules"},
	{"environment.conf", "environments/tequilaware/environment.conf"},
	{"hosts.yaml", "hieradata/hosts.yaml"},
	{"site.pp", "manifests/site.pp"},
}

func (interactor EngineInteractor) CreateConf(server domain.Server, zipFile io.Writer) {
	var Files = []domain.File{}

	packages := []domain.Package{}
	packages = server.Packages
	className := server.Hostname

	Files = append(Files, interactor.getPuppetTemplates(packages, className)...)
	Files = append(Files, interactor.getPuppetFiles()...)

	writeZip(zipFile, Files)

}

func (interactor EngineInteractor) getPuppetTemplates(packages []domain.Package, className string) []domain.File {

	var hieraClasses = []string{}
	var files = []domain.File{}
	var templates = []Template{}

	hieraClasses = append(hieraClasses, className)


	content := interactor.createPackages(packages, &hieraClasses)
	fmt.Println(hieraClasses)
	manifest := domain.Manifest{ClassName: className, Content: content}
	init := Template{
		"class",
		manifest,
		"environments/tequilaware/modules/web/manifests/init.pp",
	}
	templates = append(templates, init)
	
	init = Template{
		"hiera",
		hieraClasses,
		"hieradata/tequilaware/node/web.yaml",
	}
	templates = append(templates, init)
	
	for _, file := range templates{
		doc, error := interactor.FileWriter.WriteTemplate(file.Config, file.Name)
		if error != nil {
			fmt.Println(error)
		}
		fmt.Println(doc)
		files = append(files, domain.File{file.Path, doc})
	}
	return files
		
}

func (interactor EngineInteractor) getPuppetFiles() []domain.File {
	var files = []domain.File{}
	for _, file := range configs {
		content, err := ioutil.ReadFile(path.Join(interactor.filesPath, file.Name))
		if err != nil {
			log.Print(err)
		}
		files = append(files, domain.File{file.Path, string(content)})
	}
	return files
}

func (interactor EngineInteractor) createPackages(packages []domain.Package, hieraClasses *[]string) string {
	var manifestContent string
	for _, elem := range packages {
		if elem.Config != nil {
			switch {
			case elem.Name == "nginx":
				*hieraClasses = append(*hieraClasses, elem.Name)
				nginxConf := domain.NginxConfig{}
				json.Unmarshal(elem.Config, &nginxConf)
				doc, error := interactor.FileWriter.WriteTemplate(nginxConf, elem.Name)
				if error != nil {
					fmt.Println(error)
				}
				manifestContent += doc
			default:
				fmt.Println("Uknown config")
			}
		} else {
			doc, error := interactor.FileWriter.WriteTemplate(elem, "package")
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
