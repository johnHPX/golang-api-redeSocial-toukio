package appl

import (
	"API-RS-TOUKIO/internal/domain/users"
	"API-RS-TOUKIO/internal/infra/data/pgclient"
	"errors"
	"strings"

	"github.com/badoux/checkmail"
)

// serve para implementar os metodos do service de users
type userServiceImpl struct{}

func (userImpl *userServiceImpl) CreateUser(e *users.Entity) error {
	rep := pgclient.NewUserRepository()
	err := prepare(e, "cadastro")
	if err != nil {
		return err
	}
	return rep.CreateUser(e)
}

func (userImpl *userServiceImpl) ListALLUser() ([]users.Entity, error) {
	rep := pgclient.NewUserRepository()
	return rep.ListALLUser()
}

func (userImpl *userServiceImpl) ListByNameOrNickUsers(NameOrNick string) ([]users.Entity, error) {
	rep := pgclient.NewUserRepository()
	return rep.ListByNameOrNickUsers(NameOrNick)
}

func (userImpl *userServiceImpl) FindUser(id int64) (*users.Entity, error) {
	rep := pgclient.NewUserRepository()
	return rep.FindUser(id)
}

func (userImpl *userServiceImpl) UpdateUser(e *users.Entity) error {
	rep := pgclient.NewUserRepository()
	err := prepare(e, "atualizar")
	if err != nil {
		return err
	}
	return rep.UpdateUser(e)
}

func (userImpl *userServiceImpl) DeleteUser(id int64) error {
	rep := pgclient.NewUserRepository()
	return rep.DeleteUser(id)
}

func (userImpl *userServiceImpl) SearchforEmail(email string) (*users.Entity, error) {
	rep := pgclient.NewUserRepository()
	return rep.SearchforEmail(email)
}

// retorna todos os metodos
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

// valida os campos do request
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

// Retira os espaços em branco e verificar se há um email valido
func formatar(ent *users.Entity, etapa string) error {
	ent.Name = strings.TrimSpace(ent.Name)
	ent.Nick = strings.TrimSpace(ent.Nick)
	ent.Email = strings.TrimSpace(ent.Email)

	if etapa == "cadastro" {
		senhaComHash, erro := Hash(ent.Password)
		if erro != nil {
			return erro
		}

		ent.Password = string(senhaComHash)
	}

	return nil
}
