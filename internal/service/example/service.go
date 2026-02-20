package example

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) Handle() []byte {
	return []byte("Service worked!")
}
