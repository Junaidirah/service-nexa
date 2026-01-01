package dto

import "service-nexa/internal/models"

func ToUserResponse(u models.User) UserResponse {
	return UserResponse{
		Name:     u.Name,
		Email:    u.Email,
		Username: u.Username,
	}
}

func ToLoginResponse(u models.User, token string) LoginResponse {
	return LoginResponse{
		User:  ToUserResponse(u),
		Token: token,
	}
}

func ToUpdateProfileResponse(u models.User) UpdateProfileResponse {
	return UpdateProfileResponse{
		User: ToUserResponse(u),
	}
}

func ToChangePasswordResponse(u models.User) ChangePasswordResponse {
	return ChangePasswordResponse{
		User: ToUserResponse(u),
	}
}

func ToForgotPasswordResponse(u models.User) ForgotPasswordResponse {
	return ForgotPasswordResponse{
		User: ToUserResponse(u),
	}
}

func ToResetPasswordResponse(u models.User) ResetPasswordResponse {
	return ResetPasswordResponse{
		User: ToUserResponse(u),
	}
}

func ToUserResponseList(users []models.User) []UserResponse {
	var userResponses []UserResponse
	for _, user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}
	return userResponses
}
