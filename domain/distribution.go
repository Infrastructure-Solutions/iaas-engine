package domain

type Distribution struct {
	OS      string `json:"os"`
	Version string `json:"version"`
}

func (dist Distribution) GetOS() string {
	return dist.OS
}

func (dist Distribution) GetVersion() string {
	return dist.Version
}
