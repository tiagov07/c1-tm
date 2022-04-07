package products

type Product struct {
	Id         int     `json:"id"`
	Name       string  `json:"name"`
	Color      string  `json:"color"`
	Price      float64 `json:"price"`
	Stock      int     `json:"stock"`
	Code       string  `json:"code"`
	Posted     *bool   `json:"posted"`
	PostedDate string  `json:"postedDate"`
}

var ps []Product
var lastID int

type Repository interface {
	GetAll() ([]Product, error)
	Store(id int, name string, quantity int, price float64, code string, color string, posted bool, postedDate string) (Product, error)
	LastID() (int, error)
}

type repository struct{}

func (r *repository) GetAll() ([]Product, error) {
	return ps, nil
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}

func (r *repository) Store(id int, name string, quantity int, price float64, code string, color string, posted bool, postedDate string) (Product, error) {
	p := Product{id, name, color, price, quantity, code, &posted, postedDate}
	ps = append(ps, p)
	lastID = p.Id
	return p, nil
}

func NewRepository() Repository {
	return &repository{}
}
