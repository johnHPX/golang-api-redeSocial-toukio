package users

// funções a serem implementadas no pacote appl
type Service interface {
	// crud
	CreateUser(e *Entity) error
	ListALLUser() ([]Entity, error)
	ListByNameOrNickUsers(NameOrNick string) ([]Entity, error)
	FindUser(id int64) (*Entity, error)
	UpdateUser(e *Entity) error
	DeleteUser(id int64) error

	// login
	SearchforEmail(email string) (*Entity, error)

	// users function
	FollowUser(userID, followerID int64) error
	StopFollowing(userID, followerID int64) error
	SearchFollowers(userID int64) ([]Entity, error)
	SearchFollowing(userID int64) ([]Entity, error)

	// password securite
	SearchPassword(userID int64) (string, error)
	UpdatePassword(userID int64, password string) error
}
