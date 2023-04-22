package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"time"
)

type Config struct {
	DBMS            string
	User            string
	Password        string
	Protocol        string
	Host            string
	Port            int
	Schema          string
	Option          string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

var dbs map[string]*sql.DB

func Init(configs []Config) error {
	dbs = make(map[string]*sql.DB, 2)
	for _, config := range configs {
		connect := config.User + ":" + config.Password + "@" + config.Protocol + "(" + config.Host + ":" + strconv.Itoa(config.Port) + ")/" + config.Schema + "?" + config.Option
		db, err := sql.Open("mysql", connect)
		if err != nil {
			return err
		}
		db.SetMaxIdleConns(config.MaxIdleConns)
		db.SetMaxOpenConns(config.MaxOpenConns)
		db.SetConnMaxLifetime(config.ConnMaxLifetime)
		dbs[getKey(config.Host, config.Schema)] = db
	}
	return nil
}

func Close() error {
	for _, v := range dbs {
		err := v.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

func GetDB(host, schema string) *sql.DB {
	return dbs[getKey(host, schema)]
}

func getKey(host, schema string) string {
	return host + "/" + schema
}
