package domain

type File struct {
	Path    string `json:"path"`
	Content []byte `json:"content"`
}
