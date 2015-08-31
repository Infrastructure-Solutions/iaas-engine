package domain

type Myjson struct {
	Server Server `json:"server"`
	User   User   `json:"user"`
}

func (js Myjson) GetServer() Server {
	return js.Server
}

func (js Myjson) GetUser() User {
	return js.User
}
