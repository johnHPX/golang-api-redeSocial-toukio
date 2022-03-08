package publication

import "time"

type Entity struct {
	ID         int64
	Title      string
	Content    string
	AuthorID   int64
	AuthorNick string
	Likes      int64
	Create_at  time.Time
}
