package dto

type UserResponse struct {
	ID       string `json:"id"`
	Name     string `json:"firstName"`
	LastName string `json:"lastName"`
	Email    string `json:"email"`
	NickName string `json:"username"`
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
}
