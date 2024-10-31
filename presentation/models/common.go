package models

// All models will be wrapped in data field
// to be consistent with the JSON response
// of the API.

// Common is the wrapper for all models.
type Common struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

type Message struct {
	Message string `json:"message"`
}
