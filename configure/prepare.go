package configure

import (
	"errors"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/multierr"
)

type serviceType int

const (
	ServiceHttp serviceType = iota
	ServiceDb
	Timer
	LogFormat
)

func NewConfigure(cmd *cobra.Command, services ...serviceType) error {
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if len(services) != 0 {
		for _, item := range services {
			switch item {
			case ServiceHttp:
				configureHttpSrv(cmd)
			case ServiceDb:
				configureDb(cmd)
			case Timer:
				configureTimer(cmd)
			case LogFormat:
				logFormat(cmd)
			}
		}
	} else {
		return errors.New("ERROR - Cmd required params missing")
	}

	cmd.PreRun = func(cmd *cobra.Command, args []string) { viper.BindPFlags(cmd.Flags()) }
	return nil
}

func configureHttpSrv(cmd *cobra.Command) {
	cmd.Flags().String("http.host", "", "API host")
	cmd.Flags().Int("http.port", 0, "API port")
}

func configureDb(cmd *cobra.Command) {
	cmd.Flags().String("db.host", "", "Database server host")
	cmd.Flags().Int("db.port", 5432, "Database server port")
	cmd.Flags().String("db.name", "", "Database name")
	cmd.Flags().String("db.login", "", "Database login")
	cmd.Flags().String("db.pass", "", "Database pass")
}

func configureTimer(cmd *cobra.Command) {
	cmd.Flags().String("timer.type", "m", "s-seconds/m-minutes/h-hours - default m")
	cmd.Flags().Int64("timer.period", 1, "timer period - default 1 minute")
	cmd.Flags().Int64("limit.clean", 3, "limit last transactions - default 3")
}

func logFormat(cmd *cobra.Command) {
	cmd.Flags().String("logger.file", "STDOUT", "stdout or file")
	cmd.Flags().String("log.level", "DEBUG", "log level")
	cmd.Flags().String("log.fmt", "", "log format json or text")
}

func Validate(required ...serviceType) error {
	var validErrors error

	for _, item := range required {
		switch item {
		case ServiceHttp:
			if viper.GetString("http.host") == "" {
				validErrors = multierr.Append(validErrors, errors.New("ERROR param - http.host"))
			}
			if viper.GetInt("http.port") == 0 {
				validErrors = multierr.Append(validErrors, errors.New("ERROR param - http.port"))
			}
		case ServiceDb:
			if viper.GetString("db.host") == "" {
				validErrors = multierr.Append(validErrors, errors.New("ERROR param - db.host"))
			}
			if viper.GetInt("db.port") == 0 {
				validErrors = multierr.Append(validErrors, errors.New("ERROR param - db.port"))
			}
			if viper.GetString("db.login") == "" {
				validErrors = multierr.Append(validErrors, errors.New("ERROR param - db.login"))
			}
			if viper.GetString("db.pass") == "" {
				validErrors = multierr.Append(validErrors, errors.New("ERROR param - db.pass"))
			}
			if viper.GetString("db.name") == "" {
				validErrors = multierr.Append(validErrors, errors.New("ERROR param - db.name"))
			}
		}
	}

	if validErrors != nil {
		return validErrors
	}

	return nil
}
