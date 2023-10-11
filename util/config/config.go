package config

import (
	"log"

	"github.com/spf13/viper"
)

type (
	Config struct {
		App *Application
		Db  *Database
	}

	Application struct {
		Name          string
		Version       string
		LogOut        string
		LogFolder     string
		Port          int
		Prefork       bool
		CaseSensitive bool
		ReadTimeOut   int
		WriteTimeOut  int
	}

	Database struct {
		Url             string
		MigrationFolder string
		Host            string
		Port            int
		User            string
		Password        string
		Name            string
	}
)

func NewConfig() (config *Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	if err = viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	configuration := Config{
		App: &Application{
			Name:          viper.GetString("APP_NAME"),
			Version:       viper.GetString("APP_VERSION"),
			LogOut:        viper.GetString("APP_LOG_OUT"),
			LogFolder:     viper.GetString("APP_LOG_FOLDER"),
			Port:          viper.GetInt("APP_PORT"),
			Prefork:       viper.GetBool("APP_PREFORK"),
			CaseSensitive: viper.GetBool("APP_CASE_SENSITIVE"),
			ReadTimeOut:   viper.GetInt("APP_READ_TIMEOUT"),
			WriteTimeOut:  viper.GetInt("APP_WRITE_TIMEOUT"),
		},
		Db: &Database{
			Url:             viper.GetString("DATABASE_URL"),
			MigrationFolder: viper.GetString("MIGRATION_FOLDER"),
			Host:            viper.GetString("MYSQL_HOST"),
			Port:            viper.GetInt("MYSQL_PORT"),
			User:            viper.GetString("MYSQL_USER"),
			Password:        viper.GetString("MYSQL_PASSWORD"),
			Name:            viper.GetString("MYSQL_DBNAME"),
		},
	}
	return &configuration, err
}
