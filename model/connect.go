package model

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

type config struct {
	User     string
	Password string
	Host     string
	Port     uint16
	Name     string
}

var db *sql.DB

func init() {
	viper.SetConfigFile("config.toml")
	viper.AddConfigPath("/etc/book-mall")
	viper.AddConfigPath("$HOME/.config/book-mall")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}

	dbConfig := viper.Sub("database")

	var c config
	dbConfig.Unmarshal(&c)
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		c.User, c.Password,
		c.Host, c.Port,
		c.Name)
	db, _ = sql.Open("mysql", dsn)
}
