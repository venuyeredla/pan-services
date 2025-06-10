package models

type AuthRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type AuthResponse struct {
	Token string `json:"jwtToken"`
	Algo  string `json:"algo"`
}

type ErrorResponse struct {
	Msg    string `json:"msg"`
	Status string `json:"status"`
}
