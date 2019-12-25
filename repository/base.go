package repository

import (
	"gopkg.in/guregu/null.v3"
)

type BaseModel struct {
	Created null.Time `db:"created"`
}

type Repositories struct {
	Users        UsersRepository
	Transactions TransactionsRepository
	Sources      SourcesRepository
}
