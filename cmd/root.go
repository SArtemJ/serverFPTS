package cmd

import (
	"database/sql"
	"github.com/SArtemJ/serverFPTS/api/restapi"
	"github.com/SArtemJ/serverFPTS/api/restapi/operations"
	"github.com/SArtemJ/serverFPTS/dbdriver"
	"log"
	"strings"

	"github.com/SArtemJ/serverFPTS/configure"

	"github.com/go-openapi/loads"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "fpts",
	Short: "for payment transaction service API",
	Long:  `forPTS`,
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
		startAPI(db)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Fatalln(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "Config file path")

	err := configure.NewConfigure(rootCmd,
		configure.ServiceHttp,
		configure.ServiceDb)

	if err != nil {
		log.Fatal(err)
	}
}

func initConfig() {
	// Use config file from the flag if provided.
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
		if err := viper.ReadInConfig(); err != nil {
			logrus.WithError(err).Fatal("can't read config")
		}
	}

	viper.SetEnvPrefix("FPTS")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
}

func startAPI(db *sql.DB) {
	restapi.Db = db
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		logrus.Fatalln(err)
	}

	api := operations.NewFPTSAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	server.Host = viper.GetString("HTTP.Host")
	server.Port = viper.GetInt("HTTP.Port")

	server.ConfigureAPI()
	if err := server.Serve(); err != nil {
		logrus.Fatalln(err)
	}

	server.Shutdown()
}
