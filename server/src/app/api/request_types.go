package api

type AuthLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthRegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	FullName string `json:"fullName"`
}

func (this AuthRegisterRequest) Validate() bool {
	return this.Email != "" && this.Password != "" && this.FullName != ""
}
