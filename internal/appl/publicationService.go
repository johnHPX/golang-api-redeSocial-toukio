package appl

import (
	"API-RS-TOUKIO/internal/domain/publication"
	"API-RS-TOUKIO/internal/infra/data/pgclient"
	"errors"
	"strings"
)

// serve para implementar os metodos do service de publication
type publicationServiceImpl struct{}

func (publicationImpl *publicationServiceImpl) CreatePublication(e *publication.Entity) (int64, error) {
	rep := pgclient.NewPublicationRepository()
	err := preparePublication(e)
	if err != nil {
		return 0, err
	}
	return rep.CreatePublication(e)
}

func (publicationImpl *publicationServiceImpl) ListAllPublication(userID int64) ([]publication.Entity, error) {
	rep := pgclient.NewPublicationRepository()
	return rep.ListAllPublication(userID)
}

func (publicationImpl *publicationServiceImpl) FindByIDPublication(publicationID int64) (*publication.Entity, error) {
	rep := pgclient.NewPublicationRepository()
	return rep.FindByIDPublication(publicationID)
}

func (publicationImpl *publicationServiceImpl) UpdatePublication(publicationID int64, e *publication.Entity) error {
	rep := pgclient.NewPublicationRepository()
	err := preparePublication(e)
	if err != nil {
		return err
	}
	return rep.UpdatePublication(publicationID, e)
}

func (publicationImpl *publicationServiceImpl) DeletePublication(publicationID int64) error {
	rep := pgclient.NewPublicationRepository()
	return rep.DeletePublication(publicationID)
}

func (publicationImpl *publicationServiceImpl) ListByIDUserPublication(userID int64) ([]publication.Entity, error) {
	rep := pgclient.NewPublicationRepository()
	return rep.ListByIDUserPublication(userID)
}

func (publicationImpl *publicationServiceImpl) LikePublication(publicationID int64) error {
	rep := pgclient.NewPublicationRepository()
	return rep.LikePublication(publicationID)
}

func (publicationImpl *publicationServiceImpl) DeslikePublication(publicationID int64) error {
	rep := pgclient.NewPublicationRepository()
	return rep.DeslikePublication(publicationID)
}

// retorna todos os metodos
func NewPublicationService() publication.Service {
	return &publicationServiceImpl{}
}

// vai chamar os métodos para validar e formatar a publicação recebida
func preparePublication(ent *publication.Entity) error {
	if erro := validatePublication(ent); erro != nil {
		return erro
	}

	formatarPublication(ent)
	return nil
}

// valida os campos da requisição
func validatePublication(ent *publication.Entity) error {
	if ent.Title == "" {
		return errors.New("o titulo pe obrigatorio e não pode estar em branco")
	}

	if ent.Content == "" {
		return errors.New("O conteudo é obrigatorio e não pode estar em branco")
	}

	return nil

}

// Retira os espaços em branco
func formatarPublication(ent *publication.Entity) {
	ent.Title = strings.TrimSpace(ent.Title)
	ent.Content = strings.TrimSpace(ent.Content)
}
