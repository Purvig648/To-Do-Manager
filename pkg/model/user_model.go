package model

type SignUp struct {
	Username        string `json:"username"`
	EmailID         string `json:"email_id"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type SignIn struct {
	EmailID  string `json:"email_id"`
	Password string `json:"password"`
}

type UserResponse struct {
	Username string `json:"username"`
	EmailID  string `json:"email_id"`
}

type UserRequest struct {
	UserID uint `json:"user_id"`
}

type UserDetailsUpdate struct {
	Username string `json:"username"`
	EmailID  string `json:"email_id"`
}

type UserDetailUpdate struct {
	UpdateValue string `json:"updateValue"`
}
