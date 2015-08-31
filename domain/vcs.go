package domain

type Vcs struct {
	User_Name string `json:"user_name"`
	Url       string `json:"url"`
}

func (vcs Vcs) GetUserName() string {
	return vcs.User_Name
}

func (vcs Vcs) GetUrl() string {
	return vcs.Url
}


