package infraestructure

import (
	"bytes"
	"path"
	"text/template"
)

func WriteTemplate(conf interface{}, pack string) (string, error) {
	var templates_dir string = "infraestructure/templates"
	var conf_file string = "conf.pp.gtl"
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
