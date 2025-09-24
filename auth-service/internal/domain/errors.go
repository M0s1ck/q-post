package domain

import "errors"

var ErrNotFound = errors.New("not found")

var ErrDuplicate = errors.New("already exists")

var ErrWrongPassword = errors.New("wrong password")

var UnhandledDbError = errors.New("unhandled db error")
