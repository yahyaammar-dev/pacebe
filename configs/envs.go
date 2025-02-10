package configs

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	PublicHost             string
	Port                   string
	DBUser                 string
	DBPassword             string
	DBAddress              string
	DBName                 string
	JWTSecret              string
	JWTExpirationInSeconds int64
	STRIPE_PUBLISHABLE     string
	STRIPE_SECRET          string
}

var Envs = initConfig()

func initConfig() Config {

	return Config{
		PublicHost:             getEnv("PUBLIC_HOST", "http://localhost"),
		Port:                   getEnv("PORT", "8080"),
		DBUser:                 getEnv("DB_USER", "root"),
		DBPassword:             getEnv("DB_PASSWORD", "qw4hddqcrg123A*"),
		DBAddress:              fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
		DBName:                 getEnv("DB_NAME", "base"),
		JWTSecret:              getEnv("JWT_SECRET", "not-so-secret-now-is-it?"),
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXPIRATION_IN_SECONDS", 3600*24*7),
		STRIPE_PUBLISHABLE:     getEnv("STRIPE_PUBLISHABLE", "pk_test_51Q6ugm1DRODmqNjQ3bETSYB41sW7RFVVBnSUeqidkIFu2mk4hmwVMPX6Q8PJybS2ITcXAkGBIHD8hArvty0tv6sx00wdREvRlO"),
		STRIPE_SECRET:          getEnv("STRIPE_SECRET", "sk_test_51Q6ugm1DRODmqNjQ0ySABnnLd1Mc2cc3TCpXB49KRe90OzVax3nlflhlnia6Tz4WD58UwIa9HibQXgQxHv2xVfP300PQo8EBEN"),
	}
}

// Gets the env by key or fallbacks
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}

		return i
	}

	return fallback
}
