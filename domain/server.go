package domain

type Server struct {
	Domain       string       `json:"domain"`
	Hostname     string       `json:"hostname"`
	Provisioner  string       `json:"provisioner"`
	Distribution Distribution `json:"distribution"`
	App          App          `json:"app"`
	Packages []Package `json:"packages"`
}
