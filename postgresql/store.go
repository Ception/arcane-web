package postgresql

import (
	"fmt"
	"goarcane"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewStore(dataSourceName string) (*Store, error) {
	db, err := sqlx.Open("postgres", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] Opening Database: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("[ERROR] Connecting to Database: %w", err)
	}
	return &Store{
		AccountStore: NewAccountStore(db),
	}, nil
}

type Store struct {
	goarcane.AccountStore
}
