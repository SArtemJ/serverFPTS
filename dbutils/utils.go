package dbutils

import (
	"database/sql"
	"strings"
)

type Queryable interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
}

type Error struct {
	Err   error
	Query string
	Args  []interface{}
}

func (e Error) Error() string {
	return e.Err.Error()
}

func ExecBatch(q Queryable, batch string) error {
	requests := strings.Split(batch, ";")
	for _, request := range requests {
		request := strings.TrimSpace(request)
		if _, err := q.Exec(request); err != nil {
			return Error{Err: err, Query: request}
		}
	}
	return nil
}

func Tables(q Queryable) (tables []string, err error) {
	res, err := q.Query("SELECT tablename FROM pg_catalog.pg_tables WHERE schemaname = 'public'")
	if err != nil {
		return
	}
	for res.Next() {
		var table string
		res.Scan(&table)
		//if !strings.HasSuffix(table, "_view") {
		tables = append(tables, table)
		//}
	}
	return
}

func TableExist(q Queryable, tableName string) (bool, error) {
	tables, err := Tables(q)
	if err != nil {
		return false, err
	}
	for _, table := range tables {
		if tableName == table {
			return true, nil
		}
	}
	return false, nil
}
