package users

type Repository interface {
	CreateUser(e *Entity) error
	ListALLUser() ([]Entity, error)
	ListByNameOrNickUsers(NameOrNick string) ([]Entity, error)
	FindUser(id int64) (*Entity, error)
	// UpdateUser(id int64) error
	// DeleteUser(id int64) error
}
