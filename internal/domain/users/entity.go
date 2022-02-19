package users

import "time"

type Entity struct {
	ID        int64
	Name      string
	Nick      string
	Email     string
	Password  string
	Create_at time.Time
}
