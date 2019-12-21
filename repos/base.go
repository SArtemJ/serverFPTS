package repos

import (
	"gopkg.in/guregu/null.v3"
)

type BaseModel struct {
	Created null.Time `db:"created"`
	Updated null.Time `db:"updated"`
}

type Repositories struct {
	Users        UsersRepo
	Transactions TransactionsRepo
	Source       SourcesRepo
}
