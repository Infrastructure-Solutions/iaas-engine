package infraestructure

import (
	"bytes"
	"path"
	"text/template"
)

var templates_dir string = "infraestructure/templates"
var conf_file string = "conf.pp.gtl"

func WritePackages(conf interface{}) (string, error) {
	return writeTemplate(conf, "package")
}

func WriteConf(conf interface{}, pack string) (string, error) {
	return writeTemplate(conf, pack)
}

func WriteClass(conf interface{}) (string, error) {
	return writeTemplate(conf, "class")
}

func WriteHiera(conf interface{}) (string, error) {
	return writeTemplate(conf, "hiera")
}

func writeTemplate(conf interface{}, pack string) (string, error) {
	var doc bytes.Buffer
	t, err := template.ParseFiles(path.Join(templates_dir, pack, conf_file))
	if err != nil {
		return doc.String(), err
	}
	if e := t.Execute(&doc, conf); e != nil {
		return doc.String(), e
	}
	return doc.String(), nil
}
