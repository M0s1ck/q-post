package domain

import "errors"

var ErrNotFound = errors.New("not found")

var ErrDuplicate = errors.New("already exists")

var UnhandledDbError = errors.New("unhandled db error")

var ErrInvalidToken = errors.New("invalid access token")

var ErrInvalidDto = errors.New("invalid dto")

var ErrSelfFollow = errors.New("self follow")
