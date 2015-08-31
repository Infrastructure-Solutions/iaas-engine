package domain

type Server struct {
	Domain       string       `json:"domain"`
	Hostname     string       `json:"hostname"`
	Provisioner  string       `json:"provisioner"`
	Distribution Distribution `json:"distribution"`
	App          App          `json:"app"`
	Packages []Package `json:"packages"`
}

func (server Server) GetDomain() string {
	return server.Domain
}

func (server Server) GetHostname() string {
	return server.Hostname
}

func (server Server) GetProvisioner() string {
	return server.Provisioner
}

func (server Server) GetDistribution() Distribution {
	return server.Distribution
}

func (server Server) GetApp() App {
	return server.App
}

func (server Server) GetPackages() []Package {
	return server.Packages
}
