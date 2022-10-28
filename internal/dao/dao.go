package dao

import (
	"sync"

	"go-quick-template/internal/conf"
	"go-quick-template/internal/core"
	"go-quick-template/internal/dao/jinzhu"
	"go-quick-template/internal/dao/sakila"
	"go-quick-template/internal/dao/slonik"

	"github.com/sirupsen/logrus"
)

var (
	ds  core.DataService
	oss core.ObjectStorageService

	onceTs, onceDs, onceOss sync.Once
)

func DataService() core.DataService {
	onceDs.Do(func() {
		var v core.VersionInfo
		if conf.CfgIf("Gorm") {
			ds, v = jinzhu.NewDataService()
		} else if conf.CfgIf("Sqlx") && conf.CfgIf("MySQL") {
			ds, v = sakila.NewDataService()
		} else if conf.CfgIf("Sqlx") && (conf.CfgIf("Postgres") || conf.CfgIf("PostgreSQL")) {
			ds, v = slonik.NewDataService()
		} else {
			// default use gorm as orm for sql database
			ds, v = jinzhu.NewDataService()
		}
		logrus.Infof("use %s as data service with version %s", v.Name(), v.Version())
	})
	return ds
}

func newAuthorizationManageService() (s core.AuthorizationManageService) {
	if conf.CfgIf("Gorm") {
		s = jinzhu.NewAuthorizationManageService()
	} else if conf.CfgIf("Sqlx") && conf.CfgIf("MySQL") {
		s = sakila.NewAuthorizationManageService()
	} else if conf.CfgIf("Sqlx") && (conf.CfgIf("Postgres") || conf.CfgIf("PostgreSQL")) {
		s = slonik.NewAuthorizationManageService()
	} else {
		s = jinzhu.NewAuthorizationManageService()
	}
	return
}
