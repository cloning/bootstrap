package service

type Service struct {
	name string
}

type User struct {
	Name string
	Age  int
}

func NewService(name string) *Service {
	return &Service{name}
}

func (this *Service) GetUser() *User {
	return &User{
		this.name,
		29,
	}
}
