package dto

type UserResponse struct {
	ID        string
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Biography string `json:"biography"`
}
