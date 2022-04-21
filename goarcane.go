package goarcane

type Account struct {
	ID       int    `db:"id"`
	Username string `db:"name"`
	Password string `db:"password"`
	Email    string `db:"email"`
}

type AccountStore interface {
	Account(id int) (Account, error)
	Accounts() ([]Account, error)
	CreateAccount(t *Account) error
	UpdateAccount(t *Account) error
	DeleteAccount(id int) error
}
