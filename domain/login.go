package domain

type LoginRequest struct {
	Login    string
	Password string
}

type LoginResponse struct {
	AccessToken  string
	RefreshToken string
}
