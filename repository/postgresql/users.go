package postgresql

import (
	"github.com/SArtemJ/serverFPTS/repository"
	"github.com/didi/gendry/scanner"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	"github.com/spf13/cast"
	"gopkg.in/guregu/null.v3"
)

type UsersRepository struct {
	BaseRepository
}

func (repo *UsersRepository) initSelect(b repository.QueryBuilder, offset *int64, limit *int64, count bool) (*goqu.SelectDataset, error) {
	expression, err := builderToExpression(b, UsersTable)
	if err != nil {
		return nil, err
	}
	var q *goqu.SelectDataset
	q = newQuery().From(UsersTable).Select(
		goqu.I(UsersTable+".guid").As("guid"),
		goqu.I(UsersTable+".created").As("created"),
		goqu.I(UsersTable+".email").As("email"),
		goqu.I(UsersTable+".wallet").As("wallet"),
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

func (repo *UsersRepository) Count(b repository.UsersQueryBuilder) (int64, error) {
	q, err := repo.initSelect(b, nil, nil, true)
	if err != nil {
		return 0, err
	}
	return countFromQuery(repo.conn, UsersTable, q)
}

func (repo *UsersRepository) Collection(offset *int64, limit *int64) ([]*repository.UsersModel, error) {
	return repo.Find(nil, offset, limit)
}

func (repo *UsersRepository) Create(model *repository.UsersModel) error {
	modelMap, err := userMapFromModel(model)
	if err != nil {
		return err
	}
	_, err = insertToTable(repo.conn, UsersTable, modelMap, true)
	return err
}

func (repo *UsersRepository) Update(guid string, model *repository.UsersModel) (int64, error) {
	modelMap, err := userMapFromModel(model)
	if err != nil {
		return 0, err
	}
	return updateInTableByGUID(repo.conn, UsersTable, guid, modelMap)
}

func (repo *UsersRepository) Find(b repository.UsersQueryBuilder, offset *int64, limit *int64) ([]*repository.UsersModel, error) {
	q, err := repo.initSelect(b, offset, limit, false)
	if err != nil {
		return nil, err
	}
	sql, args, _ := q.ToSQL()

	rows, err := Query(repo.conn, sql, args...)
	if err != nil {
		return nil, err
	}

	var result []*repository.UsersModel

	dbResult, err := scanner.ScanMapClose(rows)
	if nil != err {
		return nil, err
	}

	for _, dataMap := range dbResult {
		result = append(result, userModelFromMap(dataMap))
	}

	return result, nil
}

func NewUsersRepo(conn Conn) *UsersRepository {
	return &UsersRepository{
		newBaseRepository(conn),
	}
}

func userMapFromModel(model *repository.UsersModel) (map[string]interface{}, error) {
	dataMap := make(map[string]interface{})
	if !model.GUID.IsZero() {
		dataMap["guid"] = model.GUID.String
	}
	if !model.Created.IsZero() {
		dataMap["created"] = model.Created.Time
	}
	if !model.Email.IsZero() {
		dataMap["email"] = model.Email.String
	}
	if !model.Wallet.IsZero() {
		dataMap["wallet"] = model.Wallet.Int64
	}
	return dataMap, nil
}

func userModelFromMap(data map[string]interface{}) *repository.UsersModel {
	return &repository.UsersModel{
		GUID:      null.StringFrom(cast.ToString(data["guid"])),
		BaseModel: repository.BaseModel{Created: dbValueToNullTime(data["created"])},
		Email:     null.StringFrom(cast.ToString(data["email"])),
		Wallet:    null.IntFrom(cast.ToInt64(cast.ToString(data["wallet"]))),
	}
}
