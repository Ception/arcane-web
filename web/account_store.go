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
	var tt []goarcane.Account
	err := s.Select(&tt, `SELECT * FROM accounts`)
	if err != nil {
		return []goarcane.Account{}, fmt.Errorf("[ERROR] Getting Accounts: %w", err)
	}
	return tt, nil
}

func (s *AccountStore) CreateAccount(a *goarcane.Account) error {
	err := s.Get(a, `INSERT INTO accounts VALUES ($1, $2, $3, $4) RETURNING *`,
		a.ID,
		a.Username,
		a.Password,
		a.Email)
	if err != nil {
		return fmt.Errorf("[ERROR] Creating Account: %w", err)
	}
	return nil
}

func (s *AccountStore) UpdateAccount(a *goarcane.Account) error {
	panic("not implemented") // TODO: Implement
}

func (s *AccountStore) DeleteAccount(id int) error {
	panic("not implemented") // TODO: Implement
}
