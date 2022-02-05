package server

import "time"

// ErrorResponse contains information about error
type ErrorResponse struct {
	CreatedAt time.Time
	Code      int
	Message   string
}

// SuccessResponse contains information about error
type SuccessResponse struct {
	Data interface{}
}
