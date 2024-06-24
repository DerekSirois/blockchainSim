package models

type Currency struct {
	Id    int     `db:"id"`
	Name  string  `db:"name"`
	Price float64 `db:"price"`
}
