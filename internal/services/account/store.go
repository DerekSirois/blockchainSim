package account

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Store struct {
	db *sqlx.DB
}

func (s *Store) GetById(id int) (Account, error) {
	var account Account
	err := s.db.Get(&account, "SELECT * FROM account WHERE id = $1", id)
	if err != nil {
		return account, fmt.Errorf("failed to get the account: %v", err)
	}
	return account, nil
}

func (s *Store) GetByPublicKey(key string) (Account, error) {
	var account Account
	err := s.db.Get(&account, "SELECT * FROM account WHERE publicKey = $1", key)
	if err != nil {
		return account, fmt.Errorf("failed to get the account: %v", err)
	}
	return account, nil
}

func (s *Store) Create(account Account) error {
	_, err := s.db.NamedExec("INSERT INTO account (publicKey, privateKey) VALUES (:publicKey, :privateKey)", account)
	if err != nil {
		return fmt.Errorf("failed to insert an account: %v", err)
	}
	return nil
}