package handlers

import (
	"errors"
	"github.com/SArtemJ/serverFPTS/bindata"
	"github.com/SArtemJ/serverFPTS/configureApp"
	"github.com/SArtemJ/serverFPTS/dbdriver"
	"github.com/SArtemJ/serverFPTS/dbutils"
	"github.com/SArtemJ/serverFPTS/repository"
	"github.com/SArtemJ/serverFPTS/repository/postgresql"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strings"
	"sync"
)

var initOnce sync.Once
var TestRepo repository.Repositories

func initRepositories() (err error) {
	initOnce.Do(func() {
		viper.SetEnvPrefix("FPTS")
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		viper.AutomaticEnv()
		viper.SetDefault("logger.file", "STDOUT")

		err := configureApp.Validate(
			configureApp.ServiceDb,
		)

		db, err := dbdriver.SetUpDbConnection()
		if err != nil {
			logrus.Fatal(err)
		}

		repos := postgresql.GetRepositories(db)
		TestRepo = *repos

		var c []byte
		c, err = bindata.Asset("drop.sql")
		if err != nil {
			err = errors.New("Can not clean DB: " + err.Error())
			return
		}
		err = dbutils.ExecBatch(db, string(c))
		if err != nil {
			return
		}

		c, err = bindata.Asset("init.sql")
		if err != nil {
			err = errors.New("Can not init DB: " + err.Error())
			return
		}
		err = dbutils.ExecBatch(db, string(c))
		if err != nil {
			return
		}

		err = FillRepositoriesForTest(&TestRepo)
		return
	})
	return
}
