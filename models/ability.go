package models

// Ability describes an Overwatch hero's ability.
type Ability struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Ultimate    bool   `json:"is_ultimate"`
}
