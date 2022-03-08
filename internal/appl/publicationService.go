package appl

import (
	"API-RS-TOUKIO/internal/domain/publication"
	"API-RS-TOUKIO/internal/infra/data/myclient"
)

// serve para implementar os metodos do service de publication
type publicationServiceImpl struct{}

func (publicationImpl *publicationServiceImpl) CreatePublication(e *publication.Entity) (int64, error) {
	rep := myclient.NewPublicationRepository()
	err := preparePublication(e)
	if err != nil {
		return 0, err
	}
	return rep.CreatePublication(e)
}

func (publicationImpl *publicationServiceImpl) ListAllPublication(userID int64) ([]publication.Entity, error) {
	rep := myclient.NewPublicationRepository()
	return rep.ListAllPublication(userID)
}

func (publicationImpl *publicationServiceImpl) FindByIDPublication(publicationID int64) (*publication.Entity, error) {
	rep := myclient.NewPublicationRepository()
	return rep.FindByIDPublication(publicationID)
}

func (publicationImpl *publicationServiceImpl) UpdatePublication(publicationID int64, e *publication.Entity) error {
	rep := myclient.NewPublicationRepository()
	err := preparePublication(e)
	if err != nil {
		return err
	}
	return rep.UpdatePublication(publicationID, e)
}

func (publicationImpl *publicationServiceImpl) DeletePublication(publicationID int64) error {
	rep := myclient.NewPublicationRepository()
	return rep.DeletePublication(publicationID)
}

func (publicationImpl *publicationServiceImpl) ListByIDUserPublication(userID int64) ([]publication.Entity, error) {
	rep := myclient.NewPublicationRepository()
	return rep.ListByIDUserPublication(userID)
}

func (publicationImpl *publicationServiceImpl) LikePublication(publicationID int64) error {
	db, err := myclient.Connect()
	if err != nil {
		return err
	}

	defer db.Close()

	sqlText := "update publication set likes = likes + 1 where id = ?"

	statement, err := db.Prepare(sqlText)
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.Exec(publicationID); err != nil {
		return err
	}

	return nil
}

func (publicationImpl *publicationServiceImpl) DeslikePublication(publicationID int64) error {
	db, err := myclient.Connect()
	if err != nil {
		return err
	}

	defer db.Close()

	sqlText := "update publication set likes = CASE WHEN likes > 0 THEN likes - 1 ELSE 0 END where id = ?"

	statement, err := db.Prepare(sqlText)
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.Exec(publicationID); err != nil {
		return err
	}

	return nil
}

// retorna todos os metodos
func NewPublicationService() publication.Service {
	return &publicationServiceImpl{}
}
