package goarcane

type Account struct {
	ID       int    `db:"id"`
	Username string `db:"name"`
	Password string `db:"password"`
	Email    string `db:"email"`
}

type AccountStore interface {
	GetAccount(id int) (Account, error)
	GetAccounts() ([]Account, error)
	CreateAccount(t *Account) error
	UpdateAccount(t *Account) error
	DeleteAccount(id int) error
}
