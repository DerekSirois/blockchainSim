package currency

import (
	"blockchainSim/internal/models"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Store struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) GetAll() ([]models.Currency, error) {
	currency := []models.Currency{}
	err := s.db.Select(&currency, "SELECT * FROM currency")
	if err != nil {
		return nil, fmt.Errorf("failed to get the currencies: %v", err)
	}
	return currency, nil
}

func (s *Store) GetById(id int) (models.Currency, error) {
	var currency models.Currency
	err := s.db.Get(&currency, "SELECT * FROM currency WHERE id = $1", id)
	if err != nil {
		return currency, fmt.Errorf("failed to get the currency: %v", err)
	}
	return currency, nil
}

func (s *Store) Create(currency models.Currency) error {
	_, err := s.db.NamedExec("INSERT INTO currency (name, price) VALUES (:name, :price)", currency)
	if err != nil {
		return fmt.Errorf("failed to create the currency: %v", err)
	}
	return nil
}

func (s *Store) Update(currency models.Currency) error {
	_, err := s.db.NamedExec("UPDATE currency SET price = :price WHERE id = :id", currency)
	if err != nil {
		return fmt.Errorf("failed to update the currency: %v", err)
	}
	return nil
}
