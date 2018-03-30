package model

// Person This is person type
type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

// Address This is address type
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

// ErrorResponse This is errorResponse type
type ErrorResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

type OkResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}
