package postgresql

import (
	"fmt"
	"goarcane"

	"github.com/jmoiron/sqlx"
)

type AccountStore struct {
	*sqlx.DB
}

func (s *AccountStore) Account(id int) (goarcane.Account, error) {
	var t goarcane.Account
	if err := s.Get(&t, `SELECT * FROM accounts WHERE id = $1`, id); err != nil {
		return goarcane.Account{}, fmt.Errorf("[ERROR] Getting Account: %w", err)
	}
	return t, nil
}

func (s *AccountStore) Accounts() ([]goarcane.Account, error) {
	var tt []goarcane.Account
	if err := s.Select(&tt, `SELECT * FROM accounts`); err != nil {
		return []goarcane.Account{}, fmt.Errorf("[ERROR] Getting Accounts: %w", err)
	}
	return tt, nil
}

func (s *AccountStore) CreateAccount(a *goarcane.Account) error {
	if err := s.Get(a, `INSERT INTO accounts VALUES ($1, $2, $3, $4) RETURNING *`,
		a.ID,
		a.Username,
		a.Password,
		a.Email); err != nil {
		return fmt.Errorf("[ERROR] Creating Account: %w", err)
	}
	return nil
}

func (s *AccountStore) UpdateAccount(a *goarcane.Account) error {
	if err := s.Get(a, `UPDATE accounts SET name = $1, password = $2, email = $3 WHERE id = $4 RETURNING *`,
		a.Username,
		a.Password,
		a.Email,
		a.ID); err != nil {
		return fmt.Errorf("[ERROR] Updating Account: %w", err)
	}
	return nil
}

func (s *AccountStore) DeleteAccount(id int) error {
	if _, err := s.Exec(`DELETE FROM accounts WHERE id = $1`, id); err != nil {
		return fmt.Errorf("[ERROR] Deleting Account: %w", err)
	}
	return nil
}
