package appl

import (
	"API-RS-TOUKIO/internal/domain/users"
	"API-RS-TOUKIO/internal/infra/data/myclient"
)

// serve para implementar os metodos do service de users
type userServiceImpl struct{}

func (userImpl *userServiceImpl) CreateUser(e *users.Entity) error {
	rep := myclient.NewUserRepository()
	err := prepare(e, "cadastro")
	if err != nil {
		return err
	}
	return rep.CreateUser(e)
}

func (userImpl *userServiceImpl) ListALLUser() ([]users.Entity, error) {
	rep := myclient.NewUserRepository()
	return rep.ListALLUser()
}

func (userImpl *userServiceImpl) ListByNameOrNickUsers(NameOrNick string) ([]users.Entity, error) {
	rep := myclient.NewUserRepository()
	return rep.ListByNameOrNickUsers(NameOrNick)
}

func (userImpl *userServiceImpl) FindUser(id int64) (*users.Entity, error) {
	rep := myclient.NewUserRepository()
	return rep.FindUser(id)
}

func (userImpl *userServiceImpl) UpdateUser(e *users.Entity) error {
	rep := myclient.NewUserRepository()
	err := prepare(e, "atualizar")
	if err != nil {
		return err
	}
	return rep.UpdateUser(e)
}

func (userImpl *userServiceImpl) DeleteUser(id int64) error {
	rep := myclient.NewUserRepository()
	return rep.DeleteUser(id)
}

// busca uma senha de um usuario salvo no banco atraves do seu email
func (userImpl *userServiceImpl) SearchforEmail(email string) (*users.Entity, error) {
	db, err := myclient.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	sqlText := `select
	 id, password 
	 from users 
	 where email = ?`

	row, err := db.Query(sqlText, email)
	if err != nil {
		return nil, err
	}

	var user *users.Entity
	if row.Next() {
		err := row.Scan(&user.ID, &user.Password)
		if err != nil {
			return nil, err
		}
	}

	return user, nil
}

// Permite que um usuario siga o outro
func (userImpl *userServiceImpl) FollowUser(userID, followerID int64) error {
	db, err := myclient.Connect()
	if err != nil {
		return nil
	}

	sqlText := "INSERT IGNORE INTO followers(user_id, follower_id) VALUES (?,?)"
	statement, err := db.Prepare(sqlText)
	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(userID, followerID)
	if err != nil {
		return err
	}

	return nil
}

// permite que um usuario pare de seguir outro
func (userImpl *userServiceImpl) StopFollowing(userID, followerID int64) error {
	db, err := myclient.Connect()
	if err != nil {
		return nil
	}

	sqlText := "DELETE FROM followers WHERE user_id = ? AND follower_id = ?"

	statement, err := db.Prepare(sqlText)
	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(userID, followerID)
	if err != nil {
		return err
	}

	return nil
}

// traz todos os seguidores de um usuario
func (userImpl *userServiceImpl) SearchFollowers(userID int64) ([]users.Entity, error) {
	rep := myclient.NewUserRepository()
	return rep.SearchFollowers(userID)
}

// traz todos os usuarios de um determinado usuario está seguindo
func (userImpl *userServiceImpl) SearchFollowing(userID int64) ([]users.Entity, error) {
	rep := myclient.NewUserRepository()
	return rep.SearchFollowing(userID)
}

// retorna uma senha de usuario já cadastrada no banco
func (userImpl *userServiceImpl) SearchPassword(userID int64) (string, error) {
	db, err := myclient.Connect()
	if err != nil {
		return "", nil
	}

	sqlText := `select password from users
	 where id = ?`

	rows, err := db.Query(sqlText, userID)
	if err != nil {
		return "", nil
	}

	defer rows.Close()

	var user users.Entity
	for rows.Next() {
		err := rows.Scan(&user.Password)
		if err != nil {
			return "", err
		}
	}

	return user.Password, nil
}

func (userImpl *userServiceImpl) UpdatePassword(userID int64, password string) error {
	rep := myclient.NewUserRepository()
	return rep.UpdatePassword(userID, password)
}

// retorna todos os metodos
func NewUserService() users.Service {
	return &userServiceImpl{}
}

//
// func (userImpl *userRepositoryImpl) SearchforEmail(email string) (*users.Entity, error) {

// }

// Permite que um usuario siga o outro
// func (userImpl *userRepositoryImpl) FollowUser(userID, followerID int64) error {

// }

// func (userImpl *userRepositoryImpl) StopFollowing(userID, followerID int64) error {

// }

// func (userImpl *userRepositoryImpl) SearchPassword(userID int64) (string, error) {

// }
