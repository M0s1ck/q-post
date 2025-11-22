package user

import "time"

type UserDetails struct {
	Name        *string
	Description *string
	Birthday    *time.Time
}
