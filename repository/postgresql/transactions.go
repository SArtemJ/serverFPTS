package postgresql

import (
	"github.com/SArtemJ/serverFPTS/repository"
	"github.com/didi/gendry/scanner"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	"github.com/spf13/cast"
	"gopkg.in/guregu/null.v3"
)

type TransactionsRepository struct {
	BaseRepository
}

func (repo *TransactionsRepository) initSelect(b repository.QueryBuilder, offset *int64, limit *int64, count bool) (*goqu.SelectDataset, error) {
	expression, err := builderToExpression(b, TransactionTable)
	if err != nil {
		return nil, err
	}
	var q *goqu.SelectDataset
	q = newQuery().From(TransactionTable).Select(
		goqu.I(TransactionTable+".guid").As("guid"),
		goqu.I(TransactionTable+".created").As("created"),
		goqu.I(TransactionTable+".state").As("state"),
		goqu.I(TransactionTable+".amount").As("amount"),
		goqu.I(TransactionTable+".done").As("done"),
		goqu.I(SourcesTable+".guid").As("source_guid"),
		goqu.I(SourcesTable+".type").As("source_type"),
		goqu.I(UsersTable+".guid").As("user_guid"),
	).Join(
		goqu.I(UsersTable),
		goqu.On(goqu.I(UsersTable+".guid").Eq(goqu.I(TransactionTable+".user_guid"))),
	).Join(
		goqu.I(SourcesTable),
		goqu.On(goqu.I(SourcesTable+".guid").Eq(goqu.I(TransactionTable+".source_guid"))),
	).Where(expression...).Prepared(true)

	if !count {
		order := builderToOrderExpression(b)
		if len(order) != 0 {
			q = q.Order(order...)
		}
	}

	if limit != nil {
		q = q.Limit(uint(*limit))
	}
	if offset != nil {
		q = q.Offset(uint(*offset))
	}

	return q, nil
}

func (repo *TransactionsRepository) Count(b repository.TransactionsQueryBuilder) (int64, error) {
	q, err := repo.initSelect(b, nil, nil, true)
	if err != nil {
		return 0, err
	}
	return countFromQuery(repo.conn, TransactionTable, q)
}

func (repo *TransactionsRepository) Collection(offset *int64, limit *int64) ([]*repository.TransactionModel, error) {
	return repo.Find(nil, offset, limit)
}

func (repo *TransactionsRepository) Create(model *repository.TransactionModel) error {
	modelMap, err := transactionMapFromModel(repo.conn, model)
	if err != nil {
		return err
	}
	_, err = insertToTable(repo.conn, TransactionTable, modelMap, true)
	return err
}

func (repo *TransactionsRepository) Update(guid string, model *repository.TransactionModel) (int64, error) {
	modelMap, err := transactionMapFromModel(repo.conn, model)
	if err != nil {
		return 0, err
	}
	return updateInTableByGUID(repo.conn, TransactionTable, guid, modelMap)
}

func (repo *TransactionsRepository) Find(b repository.TransactionsQueryBuilder, offset *int64, limit *int64) ([]*repository.TransactionModel, error) {
	q, err := repo.initSelect(b, offset, limit, false)
	if err != nil {
		return nil, err
	}
	sql, args, _ := q.ToSQL()

	rows, err := Query(repo.conn, sql, args...)
	if err != nil {
		return nil, err
	}

	var result []*repository.TransactionModel

	dbResult, err := scanner.ScanMapClose(rows)
	if nil != err {
		return nil, err
	}

	for _, dataMap := range dbResult {
		result = append(result, transactionModelFromMap(dataMap))
	}
	return result, nil
}

func (repo *TransactionsRepository) FindOdd(b repository.TransactionsQueryBuilder, limit *int64) ([]*repository.TransactionModel, error) {
	expression, err := builderToExpression(b, TransactionTable)
	if err != nil {
		return nil, err
	}

	nested := newQuery().From(TransactionTable).Select(
		goqu.I(TransactionTable+".id").As("id"),
		goqu.I(TransactionTable+".guid").As("guid"),
		goqu.I(TransactionTable+".state").As("state"),
		goqu.I(TransactionTable+".amount").As("amount"),
		goqu.I(TransactionTable+".done").As("done"),
		goqu.I(SourcesTable+".guid").As("source_guid"),
		goqu.I(SourcesTable+".type").As("source_type"),
		goqu.I(UsersTable+".guid").As("user_guid"),
	).Join(
		goqu.I(UsersTable),
		goqu.On(goqu.I(UsersTable+".guid").Eq(goqu.I(TransactionTable+".user_guid"))),
	).Join(
		goqu.I(SourcesTable),
		goqu.On(goqu.I(SourcesTable+".guid").Eq(goqu.I(TransactionTable+".source_guid"))),
	).Where(goqu.L(TransactionTable + ".id%2<>0"))

	var l uint
	if limit != nil {
		l = cast.ToUint(limit)
	}

	qT := newQuery().From(nested).Limit(l).Select().Where(expression...).Prepared(true)
	order := builderToOrderExpression(b)
	if len(order) != 0 {
		qT = qT.Order(order...)
	}

	sql, args, _ := qT.ToSQL()
	rows, err := Query(repo.conn, sql, args...)
	if err != nil {
		return nil, err
	}

	var result []*repository.TransactionModel

	dbResult, err := scanner.ScanMapClose(rows)
	if nil != err {
		return nil, err
	}

	for _, dataMap := range dbResult {
		result = append(result, transactionModelFromMap(dataMap))
	}

	return result, nil
}

func NewTransactionsRepo(conn Conn) *TransactionsRepository {
	return &TransactionsRepository{
		newBaseRepository(conn),
	}
}

func transactionMapFromModel(conn Conn, model *repository.TransactionModel) (map[string]interface{}, error) {
	dataMap := make(map[string]interface{})
	if !model.GUID.IsZero() {
		dataMap["guid"] = model.GUID.String
	}
	if !model.Created.IsZero() {
		dataMap["created"] = model.Created.Time
	}
	if !model.State.IsZero() {
		dataMap["state"] = model.State.String
	}
	if !model.Amount.IsZero() {
		dataMap["amount"] = model.Amount.Int64
	}
	if !model.Source.IsZero() {
		exist, err := checkExistSourceByType(conn, model.Source.String)
		if err != nil {
			return nil, err
		}
		dataMap["source_guid"] = exist
	}
	if !model.User.IsZero() {
		exist, err := checkExistUser(conn, model.User.String)
		if err != nil {
			return nil, err
		}
		dataMap["user_guid"] = exist
	}
	if !model.Done.IsZero() {
		dataMap["done"] = model.Done.Bool
	}
	return dataMap, nil
}

func transactionModelFromMap(data map[string]interface{}) *repository.TransactionModel {
	return &repository.TransactionModel{
		GUID:      null.StringFrom(cast.ToString(data["guid"])),
		BaseModel: repository.BaseModel{Created: dbValueToNullTime(data["created"])},
		State:     null.StringFrom(cast.ToString(data["state"])),
		Amount:    null.IntFrom(cast.ToInt64(cast.ToString(data["amount"]))),
		Source:    null.StringFrom(cast.ToString(data["source_type"])),
		User:      null.StringFrom(cast.ToString(data["user_guid"])),
		Done:      null.BoolFrom(cast.ToBool(data["done"])),
	}
}
