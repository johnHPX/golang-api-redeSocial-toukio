package users

type Service interface {
	CreateUser(e *Entity) error
}
