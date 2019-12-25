package cmd

import (
	"database/sql"
	"github.com/SArtemJ/serverFPTS/api/restapi"
	"github.com/SArtemJ/serverFPTS/api/restapi/operations"
	"github.com/SArtemJ/serverFPTS/calculate"
	"github.com/SArtemJ/serverFPTS/configure"
	"github.com/SArtemJ/serverFPTS/dbdriver"
	"github.com/SArtemJ/serverFPTS/repository/postgresql"
	"github.com/go-openapi/loads"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
	"time"
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

		SetLoggerSettings()

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
		configure.ServiceDb,
		configure.Timer,
		configure.LogFormat,
	)

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
	viper.AutomaticEnv()
}

func startAPI(db *sql.DB) {
	repos := postgresql.GetRepositories(db)

	timer := SetTimerLimitValues()

	var limitForBackground int64
	if viper.GetInt("limit.clean") != 0 {
		limitForBackground = viper.GetInt64("limit.clean")
	} else {
		limitForBackground = 3
	}
	calc := calculate.NewWalletCalculate(repos, limitForBackground, timer)

	restapi.Repos = repos
	restapi.Calc = calc

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		logrus.Fatalln(err)
	}

	api := operations.NewFPTSAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	server.Host = viper.GetString("HTTP.Host")
	server.Port = viper.GetInt("HTTP.Port")

	go calc.BackgroundWork()

	server.ConfigureAPI()
	if err := server.Serve(); err != nil {
		logrus.Fatalln(err)
	}

	server.Shutdown()
}

func SetTimerLimitValues() (timer *time.Ticker) {
	var intervalForBackground time.Duration

	if viper.GetInt64("timer.period") != 0 {
		intervalForBackground = time.Duration(viper.GetInt64("timer.period"))
	}

	if viper.GetString("timer.type") != "" {
		switch viper.GetString("timer.type") {
		case "s":
			timer = time.NewTicker(time.Second * intervalForBackground)
		case "m":
			timer = time.NewTicker(time.Minute * intervalForBackground)
		case "h":
			timer = time.NewTicker(time.Hour * intervalForBackground)
		}
	}
	return
}

func SetLoggerSettings() {
	logFile := viper.GetString("logger.file")
	if strings.ToUpper(logFile) != "STDOUT" {
		logFile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			logrus.WithError(err).Fatalf("can't open file")
		}

		logrus.SetOutput(logFile)
	} else {
		logrus.SetOutput(os.Stdout)
	}

	switch strings.ToLower(viper.GetString("log.fmt")) {
	case "json":
		logrus.SetFormatter(&logrus.JSONFormatter{
			PrettyPrint: false,
		})
	default:
		logrus.SetFormatter(&logrus.TextFormatter{
			ForceColors:   true,
			FullTimestamp: true,
		})
	}

	switch strings.ToLower(viper.GetString("log.level")) {
	case "fatal":
		logrus.SetLevel(logrus.FatalLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "trace":
		logrus.SetLevel(logrus.TraceLevel)
	}
}
