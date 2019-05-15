package dao

import (
	"github.com/xormplus/xorm"
	"kadmin/pkg/config"
)

type DB struct {
	Db *xorm.Engine
}

func (d DB) initDb() {
	config := config.GetGlobalConfig()
	d.Db = config.LoadDb()
}
