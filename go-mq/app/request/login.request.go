package request

type LoginRequestData struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (r *LoginRequestData) ToModel() *LoginRequestData {
	return &LoginRequestData{
		Email:    r.Email,
		Password: r.Password,
	}
}
