package configs

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/spf13/viper"
)

var cfg *config

type config struct {
	API APIConfig
	DB  DBConfig
	JWT JWTConfig
}

type APIConfig struct {
	Port      string
	RateLimit int
}

type DBConfig struct {
	Host               string
	Port               int
	User               string
	Password           string
	DBName             string
	SSLmode            string
	SetMaxIdleConns    int
	SetMaxOpenConns    int
	SetConnMaxLifetime time.Duration
}

type JWTConfig struct {
	SecretKey string
	ExpiresIn time.Duration
	Version   string
}

func InitConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("toml")

	err := viper.ReadInConfig()

	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			fmt.Printf("Error reading config file: %v\n", err)
			return
		}
	}

	switch os.Getenv("GIN_MODE") {
	case "release":
		cfg = &config{
			API: APIConfig{
				Port:      viper.GetString("api.port"),
				RateLimit: viper.GetInt("api.rate_limit"),
			},
			DB: DBConfig{
				Host:               viper.GetString("db.host"),
				Port:               viper.GetInt("db.port"),
				User:               viper.GetString("db.user"),
				Password:           viper.GetString("db.password"),
				DBName:             viper.GetString("db.dbname"),
				SSLmode:            viper.GetString("db.sslmode"),
				SetMaxIdleConns:    viper.GetInt("db.set_max_idle_conns"),
				SetMaxOpenConns:    viper.GetInt("db.set_max_open_conns"),
				SetConnMaxLifetime: viper.GetDuration("db.set_conn_max_lifetime"),
			},
			JWT: JWTConfig{
				SecretKey: viper.GetString("jwt.secret_key"),
				ExpiresIn: viper.GetDuration("jwt.expires_in"),
				Version:   viper.GetString("jwt.version"),
			},
		}
	case "test":
		cfg = &config{
			API: APIConfig{
				Port:      viper.GetString("test_api.port"),
				RateLimit: viper.GetInt("test_api.rate_limit"),
			},
			DB: DBConfig{
				Host:               viper.GetString("test_db.host"),
				Port:               viper.GetInt("test_db.port"),
				User:               viper.GetString("test_db.user"),
				Password:           viper.GetString("test_db.password"),
				DBName:             viper.GetString("test_db.dbname"),
				SSLmode:            viper.GetString("test_db.sslmode"),
				SetMaxIdleConns:    viper.GetInt("test_db.set_max_idle_conns"),
				SetMaxOpenConns:    viper.GetInt("test_db.set_max_open_conns"),
				SetConnMaxLifetime: viper.GetDuration("test_db.set_conn_max_lifetime"),
			},
			JWT: JWTConfig{
				SecretKey: viper.GetString("test_jwt.secret_key"),
				ExpiresIn: viper.GetDuration("test_jwt.expires_in"),
				Version:   viper.GetString("test_jwt.version"),
			},
		}
	case "debug":
		DBPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
		JWTExpiresIn, _ := strconv.Atoi(os.Getenv("JWT_EXPIRATION_IN"))

		cfg = &config{
			API: APIConfig{
				Port:      os.Getenv("API_PORT"),
				RateLimit: 50,
			},
			DB: DBConfig{
				Host:               os.Getenv("DB_HOST"),
				Port:               DBPort,
				User:               os.Getenv("DB_USER"),
				Password:           os.Getenv("DB_PASSWORD"),
				DBName:             os.Getenv("DB_NAME"),
				SSLmode:            "disable",
				SetMaxIdleConns:    10,
				SetMaxOpenConns:    100,
				SetConnMaxLifetime: 60,
			},
			JWT: JWTConfig{
				SecretKey: os.Getenv("JWT_SECRET_KEY"),
				ExpiresIn: time.Duration(JWTExpiresIn),
				Version:   os.Getenv("JWT_VERSION"),
			},
		}
	}
}

func GetConfig() *config {
	return cfg
}
