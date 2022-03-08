package appl

import (
	"API-RS-TOUKIO/internal/domain/users"
	"errors"
	"strings"

	"github.com/badoux/checkmail"
)

// vai chamar os métodos para validar e formatar o usuário recebido
func prepare(ent *users.Entity, etapa string) error {
	if erro := validate(ent, etapa); erro != nil {
		return erro
	}

	if erro := formate(ent, etapa); erro != nil {
		return erro
	}
	return nil
}

// valida os campos do request
func validate(ent *users.Entity, etapa string) error {
	if ent.Name == "" {
		return errors.New("o nome é obrigatório e não pode está em branco")
	}

	if ent.Nick == "" {
		return errors.New("o nick é obrigatório e não pode está em branco")
	}

	if ent.Email == "" {
		return errors.New("o email é obrigatório e não pode está em branco")
	}

	erro := checkmail.ValidateFormat(ent.Email)
	if erro != nil {
		return errors.New("o e-mail inserido é invalido")
	}

	if etapa == "cadastro" && ent.Password == "" {
		return errors.New("a senha é obrigatório e não pode está em branco")
	}

	return nil
}

// Retira os espaços em branco e verificar se há um email valido
func formate(ent *users.Entity, etapa string) error {
	ent.Name = strings.TrimSpace(ent.Name)
	ent.Nick = strings.TrimSpace(ent.Nick)
	ent.Email = strings.TrimSpace(ent.Email)

	// verifica se a requisição é um cadastro de usuario
	if etapa == "cadastro" {
		senhaComHash, erro := Hash(ent.Password) // criar uma senha com hash
		if erro != nil {
			return erro
		}

		ent.Password = string(senhaComHash) //adiciona ao entidade do response, uma senha com hash
	}

	return nil
}
