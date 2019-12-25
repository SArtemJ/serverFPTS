package postgresql

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/didi/gendry/scanner"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"gopkg.in/guregu/null.v3"
	"net"
	"time"
)

func Exec(conn Conn, sql string, args ...interface{}) (res sql.Result, err error) {
	res, err = conn.Exec(sql, args...)
	if err != nil {
		logrus.WithFields(logrus.Fields{"resource": "DB", "args": args, "sql": sql}).Tracef("Error: %s\n", err)
	} else {
		logrus.WithFields(logrus.Fields{"resource": "DB", "args": args, "sql": sql}).Trace("OK")
	}
	return res, wrapDBError(err)
}

func Query(conn Conn, sql string, args ...interface{}) (rows *sql.Rows, err error) {
	rows, err = conn.Query(sql, args...)
	if err != nil {
		logrus.WithFields(logrus.Fields{"resource": "DB", "args": args, "sql": sql}).Tracef("Error: %s\n", err)
	} else {
		logrus.WithFields(logrus.Fields{"resource": "DB", "args": args, "sql": sql}).Trace("OK")
	}
	return rows, wrapDBError(err)
}

func newQuery() *goqu.Database {
	return goqu.New("postgres", nil)
}

var ErrorSeveralID = fmt.Errorf("Insert to has returned several ID")

func insertToTableMulti(conn Conn, tableName string, data []interface{}, withoutID bool) (ids []int64, err error) {
	builder := newQuery()
	q := builder.Insert(tableName).Prepared(true).Rows(data)
	if !withoutID {
		q = q.Returning("id")
	}

	sql, args, err := q.ToSQL()
	if err != nil {
		return
	}
	if withoutID {
		res, err := Exec(conn, sql, args...)
		if err != nil {
			return ids, err
		}
		rows, err := res.RowsAffected()
		if err != nil {
			return ids, err
		}
		if int(rows) != len(data) {
			return ids, fmt.Errorf("Affected rows \"%d\" is not equal inserted \"%d\"", int(rows), len(data))
		}
		return ids, nil
	}
	rows, err := Query(conn, sql, args...)
	if err != nil {
		return
	}
	dbResult, err := scanner.ScanMapClose(rows)
	if nil != err {
		return
	}

	for _, row := range dbResult {
		if val, ok := row["id"]; ok {
			castval, err := cast.ToInt64E(val)
			if err == nil {
				ids = append(ids, castval)
			}
		}
	}

	if len(ids) != len(data) {
		return ids, fmt.Errorf("Affected rows \"%d\" is not equal inserted \"%d\"", len(ids), len(data))
	}

	return ids, nil
}

func insertToTable(conn Conn, tableName string, data interface{}, withoutID bool) (int64, error) {
	var dataMulti []interface{}
	dataMulti = append(dataMulti, data)

	res, err := insertToTableMulti(conn, tableName, dataMulti, withoutID)
	if err != nil {
		return 0, err
	}

	if withoutID {
		return 0, nil
	}
	if len(res) < 1 {
		return 0, fmt.Errorf("Unable to get the ID of the inserted row")
	} else if len(res) > 1 {
		return 0, ErrorSeveralID
	}

	return res[0], nil
}

func updateInTableByGUID(conn Conn, tableName string, guid string, data map[string]interface{}) (int64, error) {
	if len(data) == 0 {
		return 0, nil
	}
	builder := newQuery()
	sql, args, err := builder.From(tableName).Limit(1).Where(goqu.I("guid").Eq(guid)).Update().Set(data).ToSQL()
	if err != nil {
		return 0, err
	}
	res, err := Exec(conn, sql, args...)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func countFromQuery(conn Conn, table string, q *goqu.SelectDataset) (count int64, err error) {
	countField := "id"
	if table != "" {
		countField = table + "." + countField
	}
	sqlQuery, args, err := q.Select(goqu.COUNT(countField).As("count")).ToSQL()
	if err != nil {
		return 0, err
	}
	row := conn.QueryRow(sqlQuery, args...)
	err = row.Scan(&count)
	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}
	return count, nil
}

func checkExistSourceByType(conn Conn, sType string) (guid *string, err error) {
	query, args, err := goqu.From(SourcesTable).Select(
		goqu.I(SourcesTable + ".guid"),
	).Where(
		goqu.I(SourcesTable + ".type").Eq(sType),
	).ToSQL()

	if err != nil {
		return guid, err
	}

	rows, err := Query(conn, query, args...)
	if err != nil {
		return guid, err
	}

	dbResult, err := scanner.ScanMapClose(rows)
	if nil != err {
		return guid, err
	}

	switch {
	case len(dbResult) > 1:
		err = errors.New("ERROR - more than one source guid by type")
	case len(dbResult) < 1:
		err = errors.New("ERROR - no existing source guid by type")
	case len(dbResult) == 1:
		if checkValue, ok := dbResult[0]["guid"]; ok {
			res := cast.ToString(checkValue)
			guid = &res
		}
	}

	return guid, nil
}

func checkExistUser(conn Conn, uGuid string) (guid *string, err error) {
	query, args, err := goqu.From(UsersTable).Select(
		goqu.I(UsersTable + ".guid"),
	).Where(
		goqu.I(UsersTable + ".guid").Eq(uGuid),
	).ToSQL()

	if err != nil {
		return guid, err
	}

	rows, err := Query(conn, query, args...)
	if err != nil {
		return guid, err
	}

	dbResult, err := scanner.ScanMapClose(rows)
	if nil != err {
		return guid, err
	}

	switch {
	case len(dbResult) > 1:
		err = errors.New("ERROR - more than one source guid by type")
	case len(dbResult) < 1:
		err = errors.New("ERROR - no existing source guid by type")
	case len(dbResult) == 1:
		if checkValue, ok := dbResult[0]["guid"]; ok {
			res := cast.ToString(checkValue)
			guid = &res
		}
	}

	return guid, nil
}

func dbValueToNullTime(v interface{}) null.Time {
	if v == nil {
		return null.NewTime(time.Time{}, false)
	}
	return null.TimeFrom(cast.ToTime(v).Local())
}

func wrapDBError(err error) error {
	if err == nil {
		return nil
	}
	var e *net.OpError
	if errors.As(err, &e) {
		return NewConnectionError(err)
	}
	return err
}

func init() {
	goqu.SetTimeLocation(time.Local)
}

type ConnectionError struct {
	err error
}

func NewConnectionError(err error) *ConnectionError {
	return &ConnectionError{err: fmt.Errorf("Connection lost: %w", err)}
}

func (e *ConnectionError) Error() string {
	return e.err.Error()
}
