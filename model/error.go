package model

// Error is our main response model for Endpoint
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Success bool   `json:"success"`
}
