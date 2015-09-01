package domain

import (
	"encoding/json"
)

type Package struct {
	Name    string          `json:"name,omitempty"`
	Version string          `json:"version,omitempty"`
	Config  json.RawMessage `json:"config,omitempty"`
}
