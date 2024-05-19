package postgresql_database

import (
	"fmt"
	"time"
)

type Config struct {
	User                string        `json:"user" binding:"required"`
	Password            string        `json:"password" binding:"required"`
	Network             string        `json:"network" binding:"required"`
	Host                string        `json:"host" binding:"required"`
	Port                string        `json:"port" binding:"required"`
	DBName              string        `json:"dbName" binding:"required"`
	SSLMode             string        `json:"sslMode" binding:"required"`
	MaxIdleCons         int           `json:"maxIdleCons" binding:"required"`
	RequestTimeout      time.Duration `json:"requestTimeout" binding:"required"`
	MaxIdleConnDuration time.Duration `json:"maxIdleConnDuration" binding:"required"`
	DriverName          string        `json:"driverName" binding:"required"`
}

func (c *Config) ConnectionString() string {
	return fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		c.User, c.Password, c.Host, c.Port, c.DBName, c.SSLMode,
	)
}
