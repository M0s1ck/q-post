package user

type GlobalRole uint8

const (
	RoleUser GlobalRole = iota
	RoleModer
	RoleAdmin
)

var RoleIdsByNames = map[string]GlobalRole{
	"user":  RoleUser,
	"moder": RoleModer,
	"admin": RoleAdmin,
}

var RoleNamesById = map[GlobalRole]string{
	RoleUser:  "user",
	RoleModer: "moder",
	RoleAdmin: "admin",
}
