package web

import (
	"goarcane"

	"github.com/jmoiron/sqlx"
)

func NewAccountStore(db *sqlx.DB) *AccountStore {
	return &AccountStore{
		DB: db,
	}
}

type AccountStore struct {
	*sqlx.DB
}

func (s *AccountStore) Account(id int) (goarcane.Account, error) {
	panic("not implemented") // TODO: Implement
}

func (s *AccountStore) Accounts() ([]goarcane.Account, error) {
	panic("not implemented") // TODO: Implement
}

func (s *AccountStore) CreateAccount(t *goarcane.Account) error {
	panic("not implemented") // TODO: Implement
}

func (s *AccountStore) UpdateAccount(t *goarcane.Account) error {
	panic("not implemented") // TODO: Implement
}

func (s *AccountStore) DeleteAccount(id int) error {
	panic("not implemented") // TODO: Implement
}
