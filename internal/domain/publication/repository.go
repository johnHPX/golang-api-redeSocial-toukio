package publication

type Repository interface {

	// crud
	CreatePublication(e *Entity) (int64, error)
	ListAllPublication(userID int64) ([]Entity, error)
	FindByIDPublication(publicationID int64) (*Entity, error)
	UpdatePublication(publicationID int64, e *Entity) error
	DeletePublication(publicationID int64) error
	ListByIDUserPublication(userID int64) ([]Entity, error)
}
