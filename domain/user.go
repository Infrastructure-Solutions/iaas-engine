package domain

type User struct {
	Provider   Provider    `json:"provider"`
	Vcs        Vcs         `json:"vcs"`
	PublicKeys []PublicKey `json:"public_keys"`
}

func (user User) GetProvider() Provider {
	return user.Provider
}

func (user User) GetVcs() Vcs {
	return user.Vcs
}

func (user User) GetPublicKeys() []PublicKey {
	return user.PublicKeys
}
