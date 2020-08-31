package users

// UserLoginRequest is the struct to login in the application.
type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
