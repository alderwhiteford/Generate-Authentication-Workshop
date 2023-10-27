package model

type Credentials struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type ErrorResponse struct {
	Status int `json:"status"`
	Message string `json:"message"`
}
