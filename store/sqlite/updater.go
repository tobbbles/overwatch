package sqlite

import (
	"errors"
	"fmt"
	"service/models"
)

const (
	HeroUpsertQuery = `INSERT OR REPLACE INTO heros (id, name, real_name, health, armour, shield) VALUES (?, ?, ?, ?, ?, ?);`

	AbilityUpsertQuery = `INSERT OR REPLACE INTO abilities (hero, id, name, description, is_ultimate) VALUES(?, ?, ?, ?, ?)`
)

// Update creates an SQL insert for the given hero data and inserts it into the database
func (s *Store) Update(hero *models.Hero) error {
	if len(hero.Abilities) == 0 {
		return ErrMissingAbilities
	}

	fmt.Printf("%+v\n", hero)

	// Upser the hero
	if err := s.hero(hero); err != nil {
		return err
	}

	// Upsert the hero's abilities
	return s.abilities(hero)
}

func (s *Store) hero(hero *models.Hero) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(HeroUpsertQuery)
	if err != nil {
		return err
	}

	// Destructure the hero model into the Insert values
	if _, err = stmt.Exec(hero.ID, hero.Name, hero.RealName, hero.Health, hero.Armour, hero.Shield); err != nil {
		return err
	}

	stmt.Close()

	return tx.Commit()
}

func (s *Store) abilities(hero *models.Hero) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	// Recurse over abilities and upsert them into the store
	for _, ab := range hero.Abilities {
		fmt.Printf("%+v\n", ab)

		stmt, err := tx.Prepare(AbilityUpsertQuery)
		if err != nil {
			return err
		}

		if _, err := stmt.Exec(hero.ID, ab.ID, ab.Name, ab.Description, ab.Ultimate); err != nil {
			return err
		}

		stmt.Close()
	}

	return tx.Commit()
}

// Generic errors
var (
	ErrMissingAbilities = errors.New("invalid hero given - missing abilities")
)
