package dto

type UserResponse struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type LoginResponse struct {
	User  UserResponse `json:"user"`
	Token string       `json:"token"`
}

type UpdateProfileResponse struct {
	User UserResponse `json:"user"`
}

type ChangePasswordResponse struct {
	User UserResponse `json:"user"`
}

type ForgotPasswordResponse struct {
	User UserResponse `json:"user"`
}

type ResetPasswordResponse struct {
	User UserResponse `json:"user"`
}
