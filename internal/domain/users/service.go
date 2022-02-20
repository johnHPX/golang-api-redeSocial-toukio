package users

// funções a serem implementadas no pacote appl
type Service interface {
	CreateUser(e *Entity) error
	ListALLUser() ([]Entity, error)
	ListByNameOrNickUsers(NameOrNick string) ([]Entity, error)
	FindUser(id int64) (*Entity, error)
	UpdateUser(e *Entity) error
	DeleteUser(id int64) error
}
