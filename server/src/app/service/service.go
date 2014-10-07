package core

type Service struct {
	name string
}

func NewService(name string) *Service {
	return &Service{name}
}

func (this *Service) GetName() string {
	return this.name
}
