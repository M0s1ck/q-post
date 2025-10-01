package domain

type UserRole uint8

const (
	RoleUser UserRole = iota
	RoleModer
	RoleAdmin
)

var RoleIdsByNames = map[string]UserRole{
	"user":  RoleUser,
	"moder": RoleModer,
	"admin": RoleAdmin,
}

var RoleNamesById = map[UserRole]string{
	RoleUser:  "user",
	RoleModer: "moder",
	RoleAdmin: "admin",
}
