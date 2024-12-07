package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func (c *Config) connectMariadb() {
	DB, err := sql.Open("mysql", viper.GetString("MARIADB_URI"))
	if err != nil {
		panic(err.Error())
	}

	err = DB.Ping()
	if err != nil {
		panic(err.Error())
	}

	c.SqlDb = DB
	fmt.Println("SUCCESS CONNECT TO MARIADB")
}
