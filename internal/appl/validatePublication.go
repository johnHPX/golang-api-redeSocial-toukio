package appl

import (
	"API-RS-TOUKIO/internal/domain/publication"
	"errors"
	"strings"
)

// vai chamar os métodos para validar e formatar a publicação recebida
func preparePublication(ent *publication.Entity) error {
	if erro := validatePublication(ent); erro != nil {
		return erro
	}

	formatePublication(ent)
	return nil
}

// valida os campos da requisição
func validatePublication(ent *publication.Entity) error {
	if ent.Title == "" {
		return errors.New("o titulo pe obrigatorio e não pode estar em branco")
	}

	if ent.Content == "" {
		return errors.New("o conteudo é obrigatorio e não pode estar em branco")
	}

	return nil

}

// Retira os espaços em branco
func formatePublication(ent *publication.Entity) {
	ent.Title = strings.TrimSpace(ent.Title)
	ent.Content = strings.TrimSpace(ent.Content)
}
