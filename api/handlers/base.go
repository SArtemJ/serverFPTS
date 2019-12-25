package handlers

import (
	"errors"
	"github.com/SArtemJ/serverFPTS/repository"
	"github.com/SArtemJ/serverFPTS/testData"
	"go.uber.org/multierr"
	"sync"
)

var InternalServerErrorMessage = "Internal server error"

var ValidationErrorMessage = "Validation error"

var SuccessMessage = "SUCCESS"

var NotFoundMessage = "Entity not found"

var BadRequest = "Bad request"

var once sync.Once

func FillRepositoriesForTest(repos *repository.Repositories) error {
	var failed bool
	var itemsError error

	once.Do(func() {
		failed = false
		for _, item := range testData.GetTestSources() {
			if err := repos.Sources.Create(&item); err != nil {
				itemsError = multierr.Append(itemsError, errors.New("ERROR create test item - source "+item.GUID.String))
				failed = true
			}
		}

		for _, item := range testData.GetTestUsers() {
			if err := repos.Users.Create(&item); err != nil {
				itemsError = multierr.Append(itemsError, errors.New("ERROR create test item - user "+item.GUID.String))
				failed = true
			}
		}

		for _, item := range testData.GetTestTransactions() {
			if err := repos.Transactions.Create(&item); err != nil {
				itemsError = multierr.Append(itemsError, errors.New("ERROR create test item - transaction "+item.GUID.String))
				failed = true
			}
		}
	})
	if failed {
		return errors.New("filling database data failed")
	}
	return nil
}
