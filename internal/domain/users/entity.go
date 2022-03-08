package users

import "time"

// entidade de usuarios
type Entity struct {
	ID        int64
	Name      string
	Nick      string
	Email     string
	Password  string
	Create_at time.Time
}
