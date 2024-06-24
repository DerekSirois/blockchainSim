package account

import "blockchainSim/internal/models"

type AccountStore interface {
	GetById(id int) (models.Account, error)
	GetByPublicKey(key string) (models.Account, error)
	Create(account models.Account) error
}
