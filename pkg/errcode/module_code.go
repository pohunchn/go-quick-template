package errcode

var (
	UsernameHasExisted = NewError(20001, "TEST")

	NoExistUsername   = NewError(20021, "用户不存在")
	NoAdminPermission = NewError(20022, "无管理权限")
)
