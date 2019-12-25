package postgresql

import (
	"database/sql"
	"github.com/SArtemJ/serverFPTS/repository"
)

const (
	UsersTable       = "users"
	TransactionTable = "transactions"
	SourcesTable     = "sources"
)

type BaseRepository struct {
	conn Conn
}

func newBaseRepository(conn Conn) BaseRepository {
	return BaseRepository{conn: conn}
}

func GetRepositories(conn Conn) *repository.Repositories {
	return &repository.Repositories{
		Users:        NewUsersRepo(conn),
		Transactions: NewTransactionsRepo(conn),
		Sources:      NewSourcesRepo(conn),
	}
}

type Execer interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
}

type Queryer interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

type Conn interface {
	Queryer
	Execer
}
