package models

// Update describes the characteristics of an Overwatch Update
type Hero struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	RealName string `json:"real_name"`
	Health   int    `json:"health"`
	Armour   int    `json:"armour"`
	Shield   int    `json:"shield"`

	Abilities []*Ability `json:"abilities"`
}
