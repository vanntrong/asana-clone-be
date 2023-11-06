package auth

type RegisterValidation struct {
	Email    string `form:"email" json:"email" validate:"required,email"`
	Password string `form:"password" json:"password" validate:"required"`
	Name     string `form:"name" json:"name" validate:"required"`
}

type LoginValidation struct {
	Email    string `form:"email" json:"email" validate:"required,email"`
	Password string `form:"password" json:"password" validate:"required"`
}

type CheckEmailValidation struct {
	Email string `form:"email" json:"email" validate:"required,email"`
}

type LoginGoogleValidation struct {
	IdToken string `form:"id_token" json:"id_token" validate:"required"`
}

type RefreshTokenValidation struct {
	RefreshToken string `form:"refresh_token" json:"refresh_token" validate:"required"`
}
