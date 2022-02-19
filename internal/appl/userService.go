package appl

import (
	"API-RS-TOUKIO/internal/domain/users"
	"API-RS-TOUKIO/internal/infra/data/pgclient"
	"errors"
	"strings"

	"github.com/badoux/checkmail"
)

type userServiceImpl struct{}

func (s *userServiceImpl) CreateUser(e *users.Entity, etapa string) error {
	rep := pgclient.NewUserRepository()
	err := prepare(e, etapa)
	if err != nil {
		return err
	}
	return rep.CreateUser(e)
}

func (s *userServiceImpl) ListALLUser() ([]users.Entity, error) {
	rep := pgclient.NewUserRepository()
	return rep.ListALLUser()
}

func (s *userServiceImpl) ListByNameOrNickUsers(NameOrNick string) ([]users.Entity, error) {
	rep := pgclient.NewUserRepository()
	return rep.ListByNameOrNickUsers(NameOrNick)
}

func (s *userServiceImpl) FindUser(id int64) (*users.Entity, error) {
	rep := pgclient.NewUserRepository()
	return rep.FindUser(id)
}

func NewUserService() users.Service {
	return &userServiceImpl{}
}

// Preparar vai chamar os métodos para validar e formatar o usuário recebido
func prepare(ent *users.Entity, etapa string) error {
	if erro := validate(ent, etapa); erro != nil {
		return erro
	}

	if erro := formatar(ent, etapa); erro != nil {
		return erro
	}
	return nil
}

func validate(ent *users.Entity, etapa string) error {
	if ent.Name == "" {
		return errors.New("O nome é obrigatório e não pode está em branco")
	}

	if ent.Nick == "" {
		return errors.New("O nick é obrigatório e não pode está em branco")
	}

	if ent.Email == "" {
		return errors.New("O email é obrigatório e não pode está em branco")
	}

	erro := checkmail.ValidateFormat(ent.Email)
	if erro != nil {
		return errors.New("O e-mail inserido é invalido")
	}

	if etapa == "cadastro" && ent.Password == "" {
		return errors.New("A senha é obrigatório e não pode está em branco")
	}

	return nil
}

func formatar(ent *users.Entity, etapa string) error {
	ent.Name = strings.TrimSpace(ent.Name)
	ent.Nick = strings.TrimSpace(ent.Nick)
	ent.Email = strings.TrimSpace(ent.Email)

	// if etapa == "cadastro" {
	// 	senhaComHash, erro := seguranca.Hash(usuario.Senha)
	// 	if erro != nil {
	// 		return erro
	// 	}

	// 	usuario.Senha = string(senhaComHash)
	// }

	return nil
}
