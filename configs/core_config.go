package configs

import (
	"os"
	"strconv"
	"time"
)

var cfg *config

type config struct {
	API   APIConfig
	DB    DBConfig
	JWT   JWTConfig
	CACHE CacheConfig
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

type CacheConfig struct {
	Addr      string
	Port      string
	ExpiresIn time.Duration
}

type JWTConfig struct {
	ExpiresIn time.Duration
	Version   string
}

func InitConfig() {
	DBPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	JWTExpiresIn, _ := strconv.Atoi(os.Getenv("JWT_EXPIRATION_IN"))
	rateLimit, _ := strconv.Atoi(os.Getenv("API_RATE_LIMIT"))
	setMaxIdleConns, _ := strconv.Atoi(os.Getenv("DB_SET_MAX_IDLE_CONNS"))
	setMaxOpenConns, _ := strconv.Atoi(os.Getenv("DB_SET_MAX_OPEN_CONNS"))
	setConnMaxLifetime, _ := strconv.Atoi(os.Getenv("DB_SET_CONN_MAX_LIFETIME"))
	RateLimitBurst, _ := strconv.Atoi(os.Getenv("API_RATE_LIMIT_BURST"))
	CacheExpiresIn, _ := strconv.Atoi(os.Getenv("CACHE_EXPIRATION_IN"))

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
		CACHE: CacheConfig{
			Addr:      os.Getenv("CACHE_ADDR"),
			Port:      os.Getenv("CACHE_PORT"),
			ExpiresIn: time.Duration(CacheExpiresIn),
		},
		JWT: JWTConfig{
			ExpiresIn: time.Duration(JWTExpiresIn),
			Version:   os.Getenv("JWT_VERSION"),
		},
	}
}

func GetConfig() *config {
	return cfg
}
