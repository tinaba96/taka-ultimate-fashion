package config

import (
	"os"
	"strconv"
)

type DBConfig struct {
	User string
	Password string
	Host string
	Port int
	Table string
}

func GetDBConfig() DBConfig {
    port, _ := strconv.Atoi(os.Getenv("MYSQL_PORT"))
    return DBConfig{
        User: os.Getenv("MYSQL_USER"),
        Password: os.Getenv("MYSQL_PASSWORD"),
        Host: os.Getenv("MYSQL_HOST"),
        Port: port,
		Table: os.Getenv("MYSQL_DB"),
    }
}