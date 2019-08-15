package sqlite

import "service/models"

const (
	AbilityQuery             = `SELECT id, name, description, is_ultimate, hero FROM abilities WHERE id = $1;`
	AbilitiesCollectionQuery = `SELECT id, name FROM abilities;`
	HeroAbilitiesQuery       = `SELECT id, name, description, is_ultimate, hero FROM abilities JOIN heros ON abilities.hero = heros.id WHERE heros.id = $1;`
)

func (s *Store) Ability(id int) (*models.Ability, error) {

	return nil, nil
}
func (s *Store) Abilities() ([]*models.Ability, error) {

	return nil, nil
}

func (s *Store) HeroAbilities(heroID int) ([]*models.Ability, error) {

	return nil, nil
}
