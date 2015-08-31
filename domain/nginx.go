package domain

type NginxConfig struct {
	Vhost    []Vhost `json:"vhost,omitempty"`
}

type Vhost struct {
	Name    string `json:"name"`
	Path    string `json:"path"`
	Vcsrepo Vcsrepo `json:"vcsrepo,omitempty"`
}

