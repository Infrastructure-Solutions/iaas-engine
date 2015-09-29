package usecases

import (
	"encoding/json"
	"fmt"
	"github.com/iaas-engine/domain"
	"io"
)

type EngineInteractor struct {
	FileWriter FileWriter
	Git Git
}

type Git interface {
	CreateRepo(name, org string, private bool) error
	GetRepo(name string) error
	SetAuth(user, token string)
	UploadFiles(name string, files []domain.File) error
}

func NewEngineInteractor(fileWriter FileWriter, git Git) (*EngineInteractor, error) {
	interactor := &EngineInteractor{
		FileWriter: fileWriter,
		Git: git,
	}
	return interactor, nil
}

type FileWriter interface {
	WriteTemplate(conf interface{}, pack string) ([]byte, error)
	WriteZip(zipFile io.Writer, Files []domain.File)
	GetPuppetFiles() []domain.File
}

type Template struct {
	Name   string
	Config interface{}
	Path   string
}

func (interactor EngineInteractor) CreateRepo(server domain.Server, user, token string) error { 
	var Files = []domain.File{}

	interactor.Git.SetAuth(user, token)
	
	packages := []domain.Package{}
	packages = server.Packages
	className := server.Hostname

	Files = append(Files, interactor.getPuppetTemplates(packages, className)...)
	Files = append(Files, interactor.FileWriter.GetPuppetFiles()...)

	err := interactor.Git.GetRepo("infTest")

	if err != nil {
		interactor.Git.CreateRepo("infTest", "", false)
		interactor.Git.UploadFiles("infTest", Files)
		return nil
	}

	return err
}

func (interactor EngineInteractor) CreateZip(server domain.Server, zipFile io.Writer) { 
	var Files = []domain.File{}

	packages := []domain.Package{}
	packages = server.Packages
	className := server.Hostname

	Files = append(Files, interactor.getPuppetTemplates(packages, className)...)
	Files = append(Files, interactor.FileWriter.GetPuppetFiles()...)

	interactor.FileWriter.WriteZip(zipFile, Files)

}

func (interactor EngineInteractor) getPuppetTemplates(packages []domain.Package, className string) []domain.File {

	var hieraClasses = []string{}
	var files = []domain.File{}
	var templates = []Template{}

	hieraClasses = append(hieraClasses, className)


	content := interactor.createPackages(packages, &hieraClasses)

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

		files = append(files, domain.File{file.Path, doc})
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
				manifestContent += string(doc)
			default:
				fmt.Println("Uknown config")
			}
		} else {
			doc, error := interactor.FileWriter.WriteTemplate(elem, "package")
			if error != nil {
				fmt.Println(error)
			}
			manifestContent += string(doc)
		}
	}
	return manifestContent
}
