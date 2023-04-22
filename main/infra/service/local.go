package service

import (
	entsql "entgo.io/ent/dialect/sql"
	"github.com/kondroid00/sample-server-2022/main/domain/ent"
	"github.com/kondroid00/sample-server-2022/main/infra"
	"github.com/kondroid00/sample-server-2022/main/infra/db"
	pkgDB "github.com/kondroid00/sample-server-2022/package/db"
)

type Local struct {
	dbConfigMap infra.DBConfigMap
}

func (l *Local) SetDBConfigMap(dbConfigMap infra.DBConfigMap) {
	l.dbConfigMap = dbConfigMap
}

func (l *Local) GetDB() *ent.Client {
	return l.getEntClient(db.MainDB).Debug()
}

func (l *Local) GetReadDB() *ent.Client {
	return l.getEntClient(db.ReadDB).Debug()
}

func (l *Local) CloseDB() error {
	return pkgDB.Close()
}

func (l *Local) getEntClient(host db.DBHost) *ent.Client {
	db := pkgDB.GetDB(l.dbConfigMap[host].Host, l.dbConfigMap[host].Schema)
	drv := entsql.OpenDB("mysql", db)
	return ent.NewClient(ent.Driver(drv))
}
