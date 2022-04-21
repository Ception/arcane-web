package web

import "github.com/jmoiron/sqlx"

func NewAccountStore(db *sqlx.DB) *AccountStore {
	return &AccountStore{
		DB: db,
	}
}

type AccountStore struct {
	*sqlx.DB
}
