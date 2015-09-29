package domain

type RepoRequest struct {
	Owner   string `json:"owner"`
	Name    string `json:"name"`
	Private bool   `json:"private"`
	Org     string `json:"org"`
}

type MultipleFilesRequest struct {
	Author Author `json:"author"`
	Files  []File `json:"files"`
}

type Repository struct {
	Name        string `json:"name,omitempty"`
	FullName    string `json:"full_name,omitempty"`
	Description string `json:"description,omitempty"`
	Private     bool   `json:"private,omitempty"`
	HTMLURL     string `json:"html_url,omitempty"`
	CloneURL    string `json:"clone_url,omitempty"`
	SSHURL      string `json:"ssh_url,omitempty"`
}

type Author struct {
	Author  string `json:"author"`
	Message string `json:"message"`
	Branch  string `json:"branch,omitempty"`
	Email   string `json:"email"`
}

// type File struct {
// 	Path    string `json:"path"`
// 	Content []byte `json:"content"`
// }
