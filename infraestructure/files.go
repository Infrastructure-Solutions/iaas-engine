package infraestructure

import (
	"bytes"
	"path"
	"text/template"
	"io/ioutil"
	"log"
	"strings"
	"time"
	"archive/zip"
	"io"
	"github.com/iaas-engine/domain"
)

type FileWriter struct {
	templatesPath string
	filesPath  string
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

func NewFileWriter (templatesPath string, filesPath string) (*FileWriter, error) {
	fileWriter := &FileWriter{
		templatesPath: templatesPath,
		filesPath: filesPath,
	}
	return fileWriter, nil
}

func (fileWriter FileWriter) WriteTemplate(conf interface{}, pack string) ([]byte, error) {
	var conf_file string = "conf.pp.gtl"
	var doc bytes.Buffer
	
	t, err := template.ParseFiles(path.Join(fileWriter.templatesPath, pack, conf_file))
	if err != nil {
		return doc.Bytes(), err
	}
	if e := t.Execute(&doc, conf); e != nil {
		return doc.Bytes(), e
	}
	return doc.Bytes(), nil
}


func (fileWriter FileWriter) WriteZip(zipFile io.Writer, Files []domain.File) {
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
		reader := strings.NewReader(string(file.Content))
		if _, err = io.Copy(fw, reader); err != nil {
			log.Fatal(err)
		}
	}
	err := w.Close()
	if err != nil {
		log.Fatal(err)
	}
}


func (fileWriter FileWriter) GetPuppetFiles() []domain.File {
	var files = []domain.File{}
	for _, file := range configs {
		content, err := ioutil.ReadFile(path.Join(fileWriter.filesPath, file.Name))
		if err != nil {
			log.Print(err)
		}
		files = append(files, domain.File{file.Path, content})
	}
	return files
}
