package warehouse

// SuccessResponse contains successful response for warehouses
type SuccessResponse struct {
	Data interface{} `json:"data"`
}

// ErrorResponse contains information about error
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// RequestBody represents the data type that needs to be sent over request
type RequestBody struct {
	Name string
}
