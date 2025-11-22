package dto

import "github.com/google/uuid"

type ErrorResponse struct {
	Message string `json:"message" example:"Error info"`
}

type UuidOnlyResponse struct {
	Id uuid.UUID `json:"id" example:"1214a280-1162-408a-918f-5cb9300194ce"`
}
