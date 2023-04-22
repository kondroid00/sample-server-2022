package infra

import (
	"github.com/kondroid00/sample-server-2022/main/domain/ent"
	"github.com/kondroid00/sample-server-2022/main/environment"
	"github.com/kondroid00/sample-server-2022/main/infra/db"
	pkgDB "github.com/kondroid00/sample-server-2022/package/db"
	"time"
)

type (
	Service interface {
		GetDB() *ent.Client
		GetReadDB() *ent.Client
		CloseDB() error

		SetDBConfigMap(dbConfigMap DBConfigMap)
	}

	DBConfigMap map[db.DBHost]struct{ Host, Schema string }
)

func SetConfig[T Service](t T) (T, error) {
	configs := environment.GetConfig().DB.DBConfigs
	if err := initDB(configs); err != nil {
		return t, err
	}
	t.SetDBConfigMap(getDBConfigMap(configs))
	return t, nil
}

func initDB(configs []environment.DBConfig) error {
	return pkgDB.Init(convertConfigs(configs))
}

func convertConfigs(configs []environment.DBConfig) []pkgDB.Config {
	c := make([]pkgDB.Config, 0, len(configs))
	for _, v := range configs {
		c = append(c, convertConfig(v))
	}
	return c
}

func convertConfig(config environment.DBConfig) pkgDB.Config {
	return pkgDB.Config{
		DBMS:            config.DBMS,
		User:            config.User,
		Password:        config.Password,
		Protocol:        config.Protocol,
		Host:            config.Host,
		Port:            config.Port,
		Schema:          config.Schema,
		Option:          config.Option,
		MaxIdleConns:    config.MaxIdle,
		MaxOpenConns:    config.MaxOpen,
		ConnMaxLifetime: time.Duration(config.LifetimeSec) * time.Second,
	}
}

func getDBConfigMap(configs []environment.DBConfig) DBConfigMap {
	m := make(DBConfigMap, 2)
	for _, v := range configs {
		m[db.DBHost(v.Name)] = struct{ Host, Schema string }{Host: v.Host, Schema: v.Schema}
	}
	return m
}
