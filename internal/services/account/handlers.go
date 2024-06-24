package account

import (
	"blockchainSim/internal/models"
	"blockchainSim/internal/utils"
	"net/http"
)

type Handler struct {
	store AccountStore
}

func NewHandler(store AccountStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h Handler) Create(w http.ResponseWriter, r *http.Request) error {
	account := models.Account{
		PublicKey:  generateRandomKey(),
		PrivateKey: generateRandomKey(),
	}

	accountReturn := account

	privateKey, err := hashPrivateKey(account.PrivateKey)
	if err != nil {
		return err
	}
	account.PrivateKey = privateKey

	err = h.store.Create(account)
	if err != nil {
		return err
	}

	utils.Encode(w, http.StatusOK, accountReturn)
	return nil
}
