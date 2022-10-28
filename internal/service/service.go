package service

import (
	"go-quick-template/internal/conf"
	"go-quick-template/internal/core"
	"go-quick-template/internal/dao"
)

var (
	ds                 core.DataService
	oss                core.ObjectStorageService
	DisablePhoneVerify bool
)

func Initialize() {
	ds = dao.DataService()
	DisablePhoneVerify = !conf.CfgIf("Sms")
}
