package domain

type Provider struct {
	Name      string `json:"name"`
	User_Name string `json:"user_name"`
}

func (provider Provider) GetName() string {
	return provider.Name
}

func (provider Provider) GetUserName() string {
	return provider.User_Name
}
