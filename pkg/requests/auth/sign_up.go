package auth

type SignUpRequest struct {
	Username string `json:"username" binding:"required,min=3,max=32"`
	Email    string `json:"email" binding:"required,email,max=50"`
	Password string `json:"password" binding:"required,min=3,max=32"`
}
