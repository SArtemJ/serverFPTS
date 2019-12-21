package dbdriver

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/url"
)

func SetUpDbConnection() (db *sql.DB, err error) {
	logrus.WithFields(logrus.Fields{
		"action": "Db connect",
		"addr": "postgresql://" + viper.GetString("db.login") +
			":" + "*******" + "@" + viper.GetString("db.host") +
			":" + viper.GetString("db.port") + "/" +
			viper.GetString("db.name"),
	}).Debug("Try to connect to psqlDB")

	connectionStr := "postgresql://" + viper.GetString("db.login") + ":" +
		viper.GetString("db.pass") + "@" + viper.GetString("db.host") + ":" +
		viper.GetString("db.port") + "/" + viper.GetString("db.name")

	params := url.Values{}
	params.Set("sslmode", "disable")

	db, err = sql.Open("postgres", connectionStr+"?"+params.Encode())
	if err != nil {
		return
	}

	err = db.Ping()
	if err != nil {
		return
	}

	return
}
