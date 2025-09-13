package domain

import "errors"

var ErrNotFound = errors.New("not found")

var ErrDuplicate = errors.New("already exists")

var UnhandledDbError = errors.New("unhandled db error")
