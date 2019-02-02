package main

// Action ...
type Action struct {
	Payload map[string]interface{} `json:"payload"`
	Type    string                 `json:"type"`
}

// Client ...
type Client struct {
	CountryCode      string
	PhoneNumber      string
	VerificationCode string
}