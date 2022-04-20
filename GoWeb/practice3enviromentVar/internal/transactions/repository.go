package transactions

import "github/tiagov07/c1-tm/GoWeb/practice3enviromentVar/pkg/store"

type Transaction struct {
	Id              int     `json:"id"`
	TransactionCode string  `json:"transactionCode`
	Currency        string  `json:"currency"`
	Amount          float64 `json:"amount"`
	Transmitter     string  `json:"transmitter"`
	Receiver        string  `json:"receiver"`
	TransactionDate string  `json:"transactionDate"`
}

var transactions []Transaction
var lastID int

type Repository interface {
	GetAll() ([]Transaction, error)
	// Store(id int, transactionCode string, currency string, amount float64, transmitter string, receiver string, transactionDate string) (Transaction, error)
	LastID() (int, error)
}

type repository struct {
	db store.Store
}

func (r *repository) GetAll() ([]Transaction, error) {
	var transactions []Transaction
	r.db.Read(&transactions)
	return transactions, nil
}

func (r *repository) LastID() (int, error) {
	var transcations []Transaction
	if err := r.db.Read(&transcations); err != nil {
		return 0, err
	}
	if len(transcations) == 0 {
		return 0, nil
	}
	max := 0
	if len(transcations) > 0 {
		for i := range transcations {
			if transcations[i].Id > max {
				transcations[i].Id = max
			}
		}
	}
	return max, nil
}

func NewRepository(db store.Store) Repository {
	return &repository{db: db}
}
