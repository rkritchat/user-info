package config

import (
	"database/sql"
	"fmt"
	"github.com/caarlos0/env"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
)

type Cfg struct {
	DB  *sql.DB
	Env Env
}

type Env struct {
	MysqlURL      string `env:"MYSQL_URL"`
	MysqlUser     string `env:"MYSQL_USER"`
	MysqlPassword string `env:"MYSQL_PWD"`
	MysqlDBName   string `env:"MYSQL_DB_NAME"`
}

func InitConfig() *Cfg {
	//load env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}

	//parse env
	localEnv := Env{}
	err = env.Parse(&localEnv)
	if err != nil {
		log.Fatalln(err)
	}

	return &Cfg{
		DB:  initMysql(localEnv),
		Env: localEnv,
	}
}

func initMysql(env Env) *sql.DB {
	db, err := sql.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v)/%v", env.MysqlUser, env.MysqlPassword, env.MysqlURL, env.MysqlDBName))
	if err != nil {
		log.Fatalln(err)
	}
	return db
}

func (c *Cfg) Free() {
	if c.DB != nil {
		err := c.DB.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}
}
