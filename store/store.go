package store

import "service/models"

// Provider is an interface that allows you to plug multiple data stores into the application, providing it with
// hero and ability data.
type Provider interface {
	Hero(id int) (*models.Hero, error)
	Heroes() ([]*models.Hero, error)

	Ability(id int) (*models.Ability, error)
	Abilities() ([]*models.Ability, error)
}

// Updater ought to be fulfilled by the same store as the Provider, but more over the Updater is solely responsible
// for updating hero and ability data in the store.
//
// NOTE: Due to Update models containing all of their abilities - we can get away with only passing the Update.
type Updater interface {
	Update(*models.Hero) error
}
