package postgresql_config

import "fmt"

type Config struct {
	Host    string
	Port    int
	User    string
	Pass    string
	DBName  string
	SSLMode string
}

func NewConfig() *Config {
	return &Config{}
}

func (c Config) FormatDSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Pass, c.DBName, c.SSLMode)
}
