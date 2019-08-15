package sqlite

import "service/models"

const (
	AbilityQuery             = `SELECT id, name, description, is_ultimate FROM abilities WHERE id = $1;`
	AbilitiesCollectionQuery = `SELECT id, name, description, is_ultimate FROM abilities;`
	HeroAbilitiesQuery       = `SELECT abilities.id, abilities.name, abilities.description, abilities.is_ultimate FROM abilities JOIN heros ON abilities.hero = heros.id WHERE heros.id = $1;`
)

// Abilitiy will look up a specific ability by it's ID.
func (s *Store) Ability(id int) (*models.Ability, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Commit()

	stmt, err := tx.Prepare(AbilityQuery)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)

	var ability = &models.Ability{}
	if err := row.Scan(&ability.ID, &ability.Name, &ability.Description, &ability.Ultimate); err != nil {
		return nil, err
	}

	return ability, nil
}

// Abilities queries the database and returns all abilities for all heros.
func (s *Store) Abilities() ([]*models.Ability, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Commit()

	stmt, err := tx.Prepare(AbilitiesCollectionQuery)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	// Scan and grow our slice of abilities. This is gnarly, but it's the best we get when using Go's sql package.
	var abilities []*models.Ability
	for rows.Next() {
		var ability = &models.Ability{}
		if err := rows.Scan(&ability.ID, &ability.Name, &ability.Description, &ability.Ultimate); err != nil {
			return nil, err
		}

		abilities = append(abilities, ability)
	}

	return abilities, nil
}

// HeroAbilities returns all of the abilities belonging to a Hero by their ID.
func (s *Store) HeroAbilities(heroID int) ([]*models.Ability, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Commit()

	stmt, err := tx.Prepare(HeroAbilitiesQuery)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(heroID)
	if err != nil {
		return nil, err
	}

	// Scan and grow our slice of abilities. This is gnarly, but it's the best we get when using Go's sql package.
	var abilities []*models.Ability
	for rows.Next() {
		var ability = &models.Ability{}
		if err := rows.Scan(&ability.ID, &ability.Name, &ability.Description, &ability.Ultimate); err != nil {
			return nil, err
		}

		abilities = append(abilities, ability)
	}

	return abilities, nil
}
