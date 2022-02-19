package users

type Service interface {
	CreateUser(e *Entity, etapa string) error
}
