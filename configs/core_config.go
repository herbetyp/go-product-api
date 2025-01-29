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
	Port           string
	RateLimit      int
	RateLimitBurst int
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

	if os.Getenv("GIN_MODE") == "release" {
		DBPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
		JWTExpiresIn, _ := strconv.Atoi(os.Getenv("JWT_EXPIRATION_IN"))
		rateLimit, _ := strconv.Atoi(os.Getenv("API_RATE_LIMIT"))
		setMaxIdleConns, _ := strconv.Atoi(os.Getenv("DB_SET_MAX_IDLE_CONNS"))
		setMaxOpenConns, _ := strconv.Atoi(os.Getenv("DB_SET_MAX_OPEN_CONNS"))
		setConnMaxLifetime, _ := strconv.Atoi(os.Getenv("DB_SET_CONN_MAX_LIFETIME"))
		RateLimitBurst, _ := strconv.Atoi(os.Getenv("API_RATE_LIMIT_BURST"))

		cfg = &config{
			API: APIConfig{
				Port:           os.Getenv("API_PORT"),
				RateLimit:      rateLimit,
				RateLimitBurst: RateLimitBurst,
			},
			DB: DBConfig{
				Host:               os.Getenv("DB_HOST"),
				Port:               DBPort,
				User:               os.Getenv("DB_USER"),
				Password:           os.Getenv("DB_PASSWORD"),
				DBName:             os.Getenv("DB_NAME"),
				SSLmode:            os.Getenv("DB_SSLMODE"),
				SetMaxOpenConns:    setMaxOpenConns,
				SetMaxIdleConns:    setMaxIdleConns,
				SetConnMaxLifetime: time.Duration(setConnMaxLifetime),
			},
			JWT: JWTConfig{
				SecretKey: os.Getenv("JWT_SECRET_KEY"),
				ExpiresIn: time.Duration(JWTExpiresIn),
				Version:   os.Getenv("JWT_VERSION"),
			},
		}
	} else if os.Getenv("G1NM0D3") == "test" {
		cfg = &config{
			API: APIConfig{
				Port:           viper.GetString("test_api.port"),
				RateLimit:      viper.GetInt("test_api.rate_limit"),
				RateLimitBurst: viper.GetInt("test_api.rate_limit_burst"),
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
	}
}

func GetConfig() *config {
	return cfg
}
