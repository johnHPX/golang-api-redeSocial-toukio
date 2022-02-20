package users

// funções a serem implementadas, no pacote pgclient
type Repository interface {

	// metodos Crud
	CreateUser(e *Entity) error
	ListALLUser() ([]Entity, error)
	ListByNameOrNickUsers(NameOrNick string) ([]Entity, error)
	FindUser(id int64) (*Entity, error)
	UpdateUser(e *Entity) error
	DeleteUser(id int64) error

	// funções de login
	SearchforEmail(email string) (*Entity, error)
}
