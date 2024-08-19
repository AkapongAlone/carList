package requests

type LoginRequest struct {
	Username string `json:"username" binding:"required" extensions:"x-order=0"`
	Password string `json:"password" binding:"required"`
}
