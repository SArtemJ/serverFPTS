package cmd

import (
	"github.com/SArtemJ/serverFPTS/bindata"
	"github.com/SArtemJ/serverFPTS/configure"
	"github.com/SArtemJ/serverFPTS/dbdriver"
	"github.com/SArtemJ/serverFPTS/dbutils"
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

		err := configure.Validate(
			configure.ServiceHttp,
			configure.ServiceDb,
		)
		if err != nil {
			logrus.Fatal(err)
		}

		db, err := dbdriver.SetUpDbConnection()
		if err != nil {
			logrus.Fatal(err)
		}

		switch args[0] {
		case "init":
			c, err := bindata.Asset("init.sql")
			if err != nil {
				//
				os.Exit(1)
			}

			err = dbutils.ExecBatch(db, string(c))
			if err != nil {
				os.Exit(1)
			}
		case "data":
			tables, _ := dbutils.Tables(db)
			logrus.WithFields(logrus.Fields{
				"action": "Add test data in DB",
				"tables": tables,
			}).Debug("List of tables in DB")

			//TODO
		}
	},
}

func init() {
	configure.NewConfigure(dbCmd, configure.ServiceDb)
	rootCmd.AddCommand(dbCmd)
}
