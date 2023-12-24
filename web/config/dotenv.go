package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
)

func LoadEnv() (map[string]string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	env, err := godotenv.Read(filepath.Join(home, "config", ".env"))

	if err != nil {
		return nil, err
	}

	return env, nil
}

type Conn struct {
	Host string
	Port int
	User string
	Pass string
}

func (c Conn) String() string {
	return fmt.Sprintf("%s:%d, %s:%s", c.Host, c.Port, c.User, c.Pass)
}

type Config struct {
	env map[string]string
}

func Load() (Config, error) {
	m, err := LoadEnv()
	if err != nil {
		return Config{}, err
	}

	return Config{
		env: m,
	}, nil
}

func MustLoad() Config {
	c, err := Load()
	if err != nil {
		return Config{}
	}

	return c
}

func (c Config) Get(key string) string {
	return c.env[key]
}

func (c Config) GetInt(key string) int {
	v, _ := strconv.Atoi(c.env["key"])
	return v
}

func (c Config) GetJWTKey() string {
	return c.env["SIGNING_KEY"]
}

func (c Config) GetEmailConn() Conn {
	host := c.env["EMAIL_HOST"]
	port := c.GetInt("EMAIL_PORT")
	user := c.env["EMAIL_USER"]
	pass := c.env["EMAIL_PASS"]

	return Conn{
		Host: host,
		Port: port,
		User: user,
		Pass: pass,
	}
}
