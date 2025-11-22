package domain

import "errors"

var ErrNotFound = errors.New("not found")

var ErrDuplicate = errors.New("already exists")

var ErrWrongPassword = errors.New("wrong password")

var ErrInvalidToken = errors.New("token is invalid")

var ErrWeakRole = errors.New("weak role")

var UnhandledDbError = errors.New("unhandled db error")

var ErrUnexpectedApiResponse = errors.New("unexpected api response")
