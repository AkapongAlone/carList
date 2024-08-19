package responses

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
}

type TokenResponse struct {
	Status bool          `json:"status" example:"true"`
	Data   LoginResponse `json:"data"`
}