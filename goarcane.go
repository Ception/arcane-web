package goarcane

type Account struct {
	ID       int    `db:"id"`
	Username string `db:"name"`
	Password string `db:"password"`
	Email    string `db:"email"`
}

type AccountStore interface {
	User(id int) (Account, error)
	Users() ([]Account, error)
	CreateAccount(t *Account) error
	UpdateAccount(t *Thread) error
	DeleteAccount(id int) error
}
