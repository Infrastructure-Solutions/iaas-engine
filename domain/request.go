package domain

type Request struct {
	OperatingSystem string        `json:"operating_system"`
	Service         []Service     `json:"services"`
	Application     []Application `json:"apllications"`
}

type Service struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type Application struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}
