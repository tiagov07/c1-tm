package products

type service struct {
	repository Repository
}

type Service interface {
	GetAll() ([]Product, error)
	Store(name string, quantity int, price float64, code string, color string, posted bool, postedDate string) (Product, error)
}

func (s *service) Store(name string, quantity int, price float64, code string, color string, posted bool, postedDate string) (Product, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return Product{}, err
	}

	lastID++

	product, err := s.repository.Store(lastID, name, quantity, price, code, color, posted, postedDate)
	if err != nil {
		return Product{}, err
	}

	return product, nil
}

func (s *service) GetAll() ([]Product, error) {
	ps, error := s.repository.GetAll()
	if error != nil {
		return nil, error
	}
	return ps, nil
}

func NewService(r Repository) Service {
	return &service{
		repository: r}
}
