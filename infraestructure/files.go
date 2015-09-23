package infraestructure

import (
	"bytes"
	"path"
	"text/template"
)

type FileWriter struct {
	templatesPath string
}

func NewFileWriter (templatesPath string) (*FileWriter, error) {
	fileWriter := &FileWriter{ templatesPath: templatesPath }
	return fileWriter, nil
}

func (fileWriter FileWriter) WriteTemplate(conf interface{}, pack string) (string, error) {
	var conf_file string = "conf.pp.gtl"
	var doc bytes.Buffer
	
	t, err := template.ParseFiles(path.Join(fileWriter.templatesPath, pack, conf_file))
	if err != nil {
		return doc.String(), err
	}
	if e := t.Execute(&doc, conf); e != nil {
		return doc.String(), e
	}
	return doc.String(), nil
}
