package domain

type LoginRequest struct {
	Email             string `json:"email" validate:"required,email"`
	Password          string `json:"password" validate:"required"`
	NotificationToken string `json:"notification_token"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	RoleId      string `json:"role_id,omitempty"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}