package config

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
)

type Config struct {
	Server   ServerConfig
	Logger   LoggerConfig
	Postgres PostgresConfig
	Redis    RedisConfig
	Password PasswordConfig
	Cors     CorsConfig
	Otp      OtpConfig
}

type ServerConfig struct {
	Port    string
	RunMode string
}

type LoggerConfig struct {
	FilePath string
	Encoding string
	Level    string
	Logger   string
}

type PostgresConfig struct {
	Host            string
	Port            int
	User            string
	Password        string
	Dbname          string
	SSLMode         string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

type RedisConfig struct {
	Host               string
	Port               int
	User               string
	Password           string
	Db                 string
	DialTimeout        time.Duration
	ReadTimeout        time.Duration
	WriteTimeout       time.Duration
	IdleCheckFrequency time.Duration
	PoolSize           int
	PoolTimeout        time.Duration
}

type PasswordConfig struct {
	IncludeChars     bool
	IncludeDigits    bool
	MinLength        int
	MaxLength        int
	IncludeUppercase bool
	IncludeLowercase bool
}

type CorsConfig struct {
	AllowOrigins string
}

type OtpConfig struct {
	ExpireTime time.Duration
	Digits     int
	Limiter    time.Duration
}

func Getconfig() *Config {
	cfgPath := getConfigPath(os.Getenv("APP_ENV"))
	//cfgPath := getConfigPath(os.Getenv("docker"))
	v, err := loadConfig(cfgPath, "yml")
	if err != nil {
		log.Fatalf("error in load %v:", err)
	}

	cfg, err := parseConfig(v)
	if err != nil {
		log.Fatalf("error in parse %v:", err)
	}

	return cfg
}

func parseConfig(v *viper.Viper) (*Config, error) {
	var cfg Config
	err := v.Unmarshal(&cfg)
	if err != nil {
		fmt.Printf("unable to decode into config struct, %v", err)
		return nil, err
	}
	return &cfg, nil
}

func loadConfig(filename string, filetype string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigType(filetype)
	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			fmt.Printf("unable to read config struct, %v", err)
			return nil, errors.New("config file not found")
		}
		return nil, err
	}
	return v, nil
}

func getConfigPath(env string) string {
	if env == "docker" {
		return "../config/config-docker"
	} else if env == "production" {
		return "../config/config-production"
	} else {
		//return "/config/config-development"
		return "../config/config-development"

	}
}
