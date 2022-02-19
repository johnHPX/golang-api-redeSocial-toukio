package users

type Repository interface {
	CreateUser(e *Entity) error
	// ListALL() ([]Entity, error)
	// ListUsers(NameOrNick string) ([]Entity, error)
	// FindUser(id int64) (*Entity, error)
	// UpdateUser(id int64) error
	// DeleteUser(id int64) error
}
