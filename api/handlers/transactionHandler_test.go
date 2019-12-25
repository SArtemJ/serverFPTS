package handlers

import (
	"github.com/SArtemJ/serverFPTS/api/models"
	"github.com/SArtemJ/serverFPTS/api/restapi/operations"
	"github.com/SArtemJ/serverFPTS/repository"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type TransactionPostSuite struct {
	suite.Suite
	params operations.PostTransactionParams
}

var testH = NewTransactionHandlers(&TestRepo)

func (s *TransactionPostSuite) SetupTest() {
	if err := initRepositories(); err != nil {
		s.Suite.Fail(err.Error())
	}

	s.params = operations.NewPostTransactionParams()
}

func (s *TransactionPostSuite) TestOK() {
	r := s.Require()

	//unique guid
	uuid := strfmt.UUID("17013f37-05cd-4eea-a262-c2762c4e494f")
	uuidUser := strfmt.UUID("c99cec6c-7a34-4941-a988-33b52ca5c3ec")
	win := "win"
	amount := "10.17"

	paramsBody := new(models.PostTransactionParamsBody)
	paramsBody.TransactionID = &uuid
	paramsBody.UserGUID = &uuidUser
	paramsBody.State = &win
	paramsBody.Amount = &amount

	s.params.SourceType = "game"
	s.params.Body = paramsBody

	response := testH.NewPostTransaction(s.params)
	r.IsType(&operations.PostTransactionOK{}, response)
	payload := response.(*operations.PostTransactionOK).Payload

	checkResponseServiceFields(r, payload.SuccessData)

	expectedResponseObject := new(models.ResponseObject)
	expectedResponseObject.Amount = "10.17"
	expectedResponseObject.State = win
	expectedResponseObject.UserGUID = &uuidUser
	expectedResponseObject.Wallet = "11.18"

	userQ := repository.NewUserQuery()
	userQ.GUID(uuidUser.String())
	userInfo, err := TestRepo.Users.Find(userQ, nil, nil)
	r.NoError(err)
	r.Equal(len(userInfo), 1, "From DB user should be one user on unique guid %v")

	transactionQ := repository.NewTransactionQuery()
	transactionQ.GUID(uuid.String())
	transactionInfo, err := TestRepo.Transactions.Find(transactionQ, nil, nil)
	r.NoError(err)
	r.Equal(len(transactionInfo), 1, "In Db must be saved only transaction with unique guid %v")
	r.Equal(transactionInfo[0].Amount.Int64, int64(1017), "In Db value wallet and amount must be store in Int %v")

	//r.ElementsMatch(payload.Data.ResponseObject, expectedResponseObject, "Expected response no equal after request %v")
}

func checkResponseServiceFields(r *require.Assertions, payload interface{}) {
	switch p := payload.(type) {
	case models.SuccessData:
		r.Equal(*p.Message, SuccessMessage)
	}
}

func TestPostTransaction(t *testing.T) {
	suite.Run(t, new(TransactionPostSuite))
}
