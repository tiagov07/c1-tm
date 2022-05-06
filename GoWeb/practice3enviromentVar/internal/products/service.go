package products

type Service interface {
	GetAll() ([]Product, error)
	Store(name, productType string, count int, price float64) (Product, error)
	Update(id int, name, productType string, count int, price float64) (Product, error)
	UpdateName(id int, name string) (Product, error)
	Delete(id int) error
	GetByName(name string) (Product, error)
}

type service struct {
	repository Repository
}

func (s *service) GetAll() ([]Product, error) {
	ps, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return ps, nil
}

func (s *service) GetByName(name string) (Product, error) {
	product, err := s.repository.GetByName(name)
	if err != nil {
		return Product{}, err
	}
	return product, nil
}

func (s *service) Store(name, productType string, count int, price float64) (Product, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return Product{}, err
	}
	lastID++
	p := Product{0, name, productType, count, price}
	product, err := s.repository.Store(p)
	if err != nil {
		return Product{}, err
	}
	return product, nil
}

func (s *service) Update(id int, name, productType string, count int, price float64) (Product, error) {
	return s.repository.Update(id, name, productType, count, price)
}

func (s *service) UpdateName(id int, name string) (Product, error) {
	return s.repository.UpdateName(id, name)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}
