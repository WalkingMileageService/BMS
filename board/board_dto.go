package board

import "time"

type Board struct {
	Title   string
	Content string
	UserId  string
	Created time.Time
	Updated time.Time
}
