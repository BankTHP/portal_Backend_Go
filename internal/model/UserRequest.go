package model

type CreateUserRequest struct {
	UserId     string `json:"user_id" validate:"required"`
	Name       string `json:"name" validate:"required"`
	Username   string `json:"username" validate:"required"`
	GivenName  string `json:"given_name"`
	FamilyName string `json:"family_name"`
	Email      string `json:"email" validate:"required,email"`
}

type UpdateUserRequest struct {
	UserId     string `json:"user_id" validate:"required"`
	Name       string `json:"name"`
	GivenName  string `json:"given_name"`
	FamilyName string `json:"family_name"`
	Email      string `json:"email" validate:"email"`
}
