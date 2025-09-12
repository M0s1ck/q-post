package domain

import "errors"

var ErrorNotFound = errors.New("not found")

var UnhandledDbError = errors.New("unhandled db error")
