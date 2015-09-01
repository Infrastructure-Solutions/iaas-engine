package domain

type Vcsrepo struct {
	Path    string `json:"path"`
	Version string `json:"version"`
	Source  string `json:"source"`
}
