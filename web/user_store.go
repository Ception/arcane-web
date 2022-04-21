package web

import "github.com/jmoiron/sqlx"

func NewUserStore(db *sqlx.DB) *UserStore {

}

type UserStore struct {
	*sqlx.DB
}
