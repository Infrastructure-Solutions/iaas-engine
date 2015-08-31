package domain

type PublicKey struct {
	Title string `json:"title"`
	Key   string `json:"key"`
}

func (publickey PublicKey) GetTitle() string {
	return publickey.Title
}

func (publickey PublicKey) GetKey() string {
	return publickey.Key
}
