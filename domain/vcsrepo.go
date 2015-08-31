package domain

type Vcsrepo struct {
	Path    string `json:"path"`
	Version string `json:"version"`
	Source  string `json:"source"`
}

func (vcsrepo Vcsrepo) GetPath() string {
	return vcsrepo.Path
}

func (vcsrepo Vcsrepo) GetVersion() string {
	return vcsrepo.Version
}

func (vcsrepo Vcsrepo) GetSource() string {
	return vcsrepo.Source
}
