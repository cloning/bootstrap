package user

type EmailAlreadyExists struct {
}

func (this EmailAlreadyExists) Error() string {
	return "Email already exists"
}
