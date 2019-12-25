package cmd

import (
	"github.com/SArtemJ/serverFPTS/api/handlers"
	"github.com/SArtemJ/serverFPTS/bindata"
	"github.com/SArtemJ/serverFPTS/configureApp"
	"github.com/SArtemJ/serverFPTS/dbdriver"
	"github.com/SArtemJ/serverFPTS/dbutils"
	"github.com/SArtemJ/serverFPTS/repository/postgresql"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "init or test data",
	Args: func(cmd *cobra.Command, args []string) error {
		return cobra.OnlyValidArgs(cmd, []string{"init", "data"})
	},
	Run: func(cmd *cobra.Command, args []string) {

		err := configureApp.Validate(
			configureApp.ServiceHttp,
			configureApp.ServiceDb,
		)

		db, err := dbdriver.SetUpDbConnection()
		if err != nil {
			logrus.Fatal(err)
		}

		switch args[0] {
		case "init":
			c, err := bindata.Asset("init.sql")
			if err != nil {
				logrus.Fatal(err)
				os.Exit(1)
			}

			err = dbutils.ExecBatch(db, string(c))
			if err != nil {
				logrus.Fatal(err)
				os.Exit(1)
			}
		case "drop":
			c, err := bindata.Asset("drop.sql")
			if err != nil {
				logrus.Fatal(err)
				os.Exit(1)
			}

			err = dbutils.ExecBatch(db, string(c))
			if err != nil {
				logrus.Fatal(err)
				os.Exit(1)
			}
		case "data":
			tables, _ := dbutils.Tables(db)
			logrus.WithFields(logrus.Fields{
				"action": "Add test data in DB",
				"tables": tables,
			}).Debug("List of tables in DB")

			repos := postgresql.GetRepositories(db)
			err := handlers.FillRepositoriesForTest(repos)
			if err != nil {
				logrus.Fatal(err)
			}
		}
	},
}

func init() {
	configureApp.NewConfigure(dbCmd, configureApp.ServiceDb, configureApp.LogFormat)
	rootCmd.AddCommand(dbCmd)
}
