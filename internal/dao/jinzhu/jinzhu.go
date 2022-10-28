// Core service implement base gorm+mysql/postgresql/sqlite3.
// Jinzhu is the primary developer of gorm so use his name as
// pakcage name as a saluter.

package jinzhu

import (
	"go-quick-template/internal/conf"
	"go-quick-template/internal/core"

	"github.com/Masterminds/semver/v3"
	"github.com/sirupsen/logrus"
)

var (
	_ core.DataService = (*dataServant)(nil)
	_ core.VersionInfo = (*dataServant)(nil)
)

type dataServant struct {
	core.UserManageService
}

func NewDataService() (core.DataService, core.VersionInfo) {
	// initialize CacheIndex if needed
	var (
		v core.VersionInfo
	)
	db := conf.MustGormDB()

	logrus.Infof("use %s as cache index service by version: %s", v.Name(), v.Version())

	ds := &dataServant{
		UserManageService: newUserManageService(db),
	}
	return ds, ds
}

func NewAuthorizationManageService() core.AuthorizationManageService {
	return &authorizationManageServant{
		db: conf.MustGormDB(),
	}
}

func (s *dataServant) Name() string {
	return "Gorm"
}

func (s *dataServant) Version() *semver.Version {
	return semver.MustParse("v0.1.0")
}
