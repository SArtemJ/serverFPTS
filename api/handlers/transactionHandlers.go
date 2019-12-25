package handlers

import (
	"errors"
	"github.com/SArtemJ/serverFPTS/api/models"
	"github.com/SArtemJ/serverFPTS/api/restapi/operations"
	"github.com/SArtemJ/serverFPTS/calculator"
	"github.com/SArtemJ/serverFPTS/repository"
	"github.com/SArtemJ/serverFPTS/utils"
	"github.com/go-openapi/runtime/middleware"
	"github.com/spf13/cast"
	"gopkg.in/guregu/null.v3"
	"strconv"
	"time"
)

type TransactionHandlers struct {
	repos      *repository.Repositories
	calculator *calculator.WalletCalculator
}

func NewTransactionHandlers(repos *repository.Repositories) *TransactionHandlers {
	return &TransactionHandlers{
		repos: repos,
	}
}

var Action = "Post Transaction Request"

func (th *TransactionHandlers) NewPostTransaction(params operations.PostTransactionParams) middleware.Responder {

	th.calculator = calculator.NewWalletCalculator()
	//check source type
	sourceBuilder := repository.NewSourceQuery()
	if params.SourceType != "" {
		sourceBuilder.Type(params.SourceType)
	}

	_, err := th.repos.Sources.Find(sourceBuilder, nil, nil)
	if err != nil {
		return NotFoundPayload(err)
	}

	//check transactionID
	tBuilder := repository.NewTransactionQuery()
	if params.Body.TransactionID != nil {
		tBuilder.UserGUID(params.Body.TransactionID.String())
		existT, err := th.repos.Transactions.Find(tBuilder, nil, nil)
		if err != nil {
			return NotFoundPayload(errors.New("ERROR - Can't check - transactionID"))
		}
		if len(existT) > 0 {
			return BadRequestPayload(errors.New("ERROR - transactionID already exist"))
		}
	} else {
		return BadRequestPayload(errors.New("ERROR - wrong transactionID incoming value"))
	}

	//check userGUID
	userBuilder := repository.NewUserQuery()
	var userFromDb repository.UsersModel
	if params.Body.UserGUID != nil {
		userBuilder.GUID(params.Body.UserGUID.String())
		users, err := th.repos.Users.Find(userBuilder, nil, nil)
		if err != nil {
			return NotFoundPayload(err)
		}
		if len(users) > 1 {
			return BadRequestPayload(errors.New("ERROR - on this userGUID more than one rows"))
		}
		userFromDb = *users[0]
	} else {
		err := errors.New("ERROR - wrong userGUID in request")
		BadRequestPayload(err)
	}

	//if all ok calculator wallet
	transactionIn, err := WebTransactionToDb(params.Body.TransactionObject, null.StringFrom(params.SourceType))
	if err != nil || transactionIn == nil {
		return InternalServerErrorPayload(err)
	}

	userWallet, err := th.calculator.Calculate(userFromDb.Wallet.Int64, transactionIn.Amount.Int64, transactionIn.State.String)
	if err != nil {
		return InternalServerErrorPayload(err)
	}

	userFromDb.Wallet = null.IntFrom(userWallet)
	_, err = th.repos.Users.Update(params.Body.UserGUID.String(), &userFromDb)
	if err != nil {
		return InternalServerErrorPayload(err)
	}

	err = th.repos.Transactions.Create(transactionIn)
	if err != nil {
		return InternalServerErrorPayload(err)
	}

	//response body
	payload := new(models.PostTransactionOKBody)

	data := new(models.PostTransactionOKBodyAllOf1Data)
	data.UserGUID = params.Body.UserGUID
	w := utils.AmountWalletRuleExtract(userFromDb.Wallet.Int64)
	data.Wallet = cast.ToString(w)
	data.Amount = cast.ToString(utils.AmountWalletRuleExtract(transactionIn.Amount.Int64))
	data.State = transactionIn.State.String

	payload.Message = &SuccessMessage
	payload.Data = data
	return operations.NewPostTransactionOK().WithPayload(payload)
	//TODO hTest
}

func WebTransactionToDb(object models.TransactionObject, sourceName null.String) (*repository.TransactionModel, error) {
	var intAmount int64
	var err error

	if object.Amount != nil {
		parseF, err := strconv.ParseFloat(*object.Amount, 64)
		if err != nil {
			return nil, err
		}
		intAmount = utils.AmountWalletRuleSave(parseF)
	} else {
		return nil, errors.New("ERROR - wrong amount value in transaction")
	}

	return &repository.TransactionModel{
		BaseModel: repository.BaseModel{Created: null.TimeFrom(time.Now())},
		GUID:      null.StringFrom(object.TransactionID.String()),
		State:     null.StringFromPtr(object.State),
		Amount:    null.IntFrom(intAmount),
		Source:    sourceName,
		User:      null.StringFrom(object.UserGUID.String()),
	}, err
}

func NotFoundPayload(err error) middleware.Responder {
	payload := new(models.PostTransactionNotFoundBody)
	payload.Message = &NotFoundMessage
	m := err.Error()
	payload.Errors = &m
	return operations.NewPostTransactionNotFound().WithPayload(payload)
}

func InternalServerErrorPayload(err error) middleware.Responder {
	payload := new(models.PostTransactionInternalServerErrorBody)
	payload.Message = &NotFoundMessage
	m := err.Error()
	payload.Errors = &m
	return operations.NewPostTransactionInternalServerError().WithPayload(payload)
}

func BadRequestPayload(err error) middleware.Responder {
	payload := new(models.PostTransactionBadRequestBody)
	payload.Message = &BadRequest
	m := err.Error()
	payload.Errors = &m
	return operations.NewPostTransactionBadRequest().WithPayload(payload)
}
