package domain

type User struct {
	Provider   Provider    `json:"provider"`
	Vcs        Vcs         `json:"vcs"`
	PublicKeys []PublicKey `json:"public_keys"`
}
