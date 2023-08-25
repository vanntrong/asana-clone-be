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
