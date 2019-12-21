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
			}
		}
	} else {
		//
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

// Validate start parameters
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
