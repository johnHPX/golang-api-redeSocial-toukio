package users

// funções a serem implementadas, no pacote pgclient
type Repository interface {

	// crud
	CreateUser(e *Entity) error
	ListALLUser() ([]Entity, error)
	ListByNameOrNickUsers(NameOrNick string) ([]Entity, error)
	FindUser(id int64) (*Entity, error)
	UpdateUser(e *Entity) error
	DeleteUser(id int64) error
	SearchFollowers(userID int64) ([]Entity, error)
	SearchFollowing(userID int64) ([]Entity, error)
	UpdatePassword(userID int64, password string) error
}
