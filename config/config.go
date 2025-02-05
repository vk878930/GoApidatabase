package config

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)
 

type DBConfig struct {
	Host string
	Port string
	Database string
	Username string
	Password string
	Driver string
}

type APIConfig struct {
	ApiPort string
}

type TokenConfig struct {
	ApplicationName string //menampung data
	JwtSignatureKey []byte //menampung secret key dalam slice of byte untuk proses penandatanganan token
	JwtSigningMethod *jwt.SigningMethodHMAC //pointer ini merujuk pada library objek dari method penandatanganan HMAC dari library JWT
	AccessTokenLifeTime time.Duration //menmapung durasi masa berlaku token 
}

type Config struct{
	DBConfig
	APIConfig
	TokenConfig
}

func (c *Config) readConfig() error {
	c.DBConfig = DBConfig{
		Host: "localhost",
		Port: "5432",
		Database: "enigma_laundry",
		Username: "postgres",
		Password: "2001",
		Driver: "postgres",
	}

	c.APIConfig = APIConfig{
		ApiPort: "8080",
	}

	accessTokenLifeTime := time.Duration(1) * time.Hour

	c.TokenConfig = TokenConfig{
		ApplicationName: "Enigma Laundry",
		JwtSignatureKey: []byte("Enigma Laundry"),
		JwtSigningMethod: jwt.SigningMethodHS256,
		AccessTokenLifeTime: accessTokenLifeTime,
	}

	if c.Host == "" || c.Port == "" || c.Username == "" || c.Password == "" || c.ApiPort == "" {
		return fmt.Errorf("required config")
	}

	return nil
}

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cfg.readConfig()

	if err != nil {
		return nil, err
	}

	return cfg, nil
}