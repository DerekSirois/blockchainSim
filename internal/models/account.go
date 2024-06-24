package models

type Account struct {
	Id         int
	PublicKey  string `db:"publicKey"`
	PrivateKey string `db:"privateKey"`
}
