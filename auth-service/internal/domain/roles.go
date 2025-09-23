package domain

type UserRole uint8

const (
	RoleUser UserRole = iota
	RoleModer
	RoleAdmin
)
