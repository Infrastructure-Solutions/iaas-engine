package domain

import (
	"encoding/json"
)

type Package struct {
	Name    string          `json:"name,omitempty"`
	Version string          `json:"version,omitempty"`
	Config  json.RawMessage `json:"config,omitempty"`
}

func (pack Package) GetName() string {
	return pack.Name
}

func (pack Package) GetVersion() string {
	return pack.Version
}

func (pack Package) GetConfig() json.RawMessage {
	return pack.Config
}
