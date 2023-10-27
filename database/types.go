package database

type User struct {
	Id string `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

type SecretNote struct {
	Id string `json:"id,omitempty"`
	Author string `json:"author,omitempty"`
	Note string `json:"note,omitempty"`
}
