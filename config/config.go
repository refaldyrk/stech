package config

import (
	"database/sql"
	"fmt"
	"github.com/minio/minio-go/v7"
)

type Config struct {
	SqlDb       *sql.DB
	MinioClient *minio.Client
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) InitialConfig() {
	fmt.Println("INITIAL CONFIG")
	c.connectMariadb()
	c.minioConnect()
	fmt.Println("SUCCESS INITIAL CONFIG")
}
