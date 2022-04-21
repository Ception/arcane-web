package web

import (
	"fmt"
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
	var t goarcane.Account
	err := s.Get(&t, `SELECT * FROM accounts WHERE id = $1`, id)
	if err != nil {
		return goarcane.Account{}, fmt.Errorf("[ERROR] Getting Account: %w", err)
	}
	return t, nil
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
