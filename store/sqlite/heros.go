package sqlite

import "service/models"

const (
	HeroQuery            = `SELECT id, name, real_name, health, armour, shield FROM heros WHERE id = $1;`
	HerosCollectionQuery = `SELECT id, name, real_name, health, armour, shield FROM heros ORDER BY id;
`
)

// Hero provides a single hero from the store
func (s *Store) Hero(id int) (*models.Hero, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Commit()

	stmt, err := tx.Prepare(HeroQuery)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)

	var hero = &models.Hero{}
	if err := row.Scan(&hero.ID, &hero.Name, &hero.RealName, &hero.Health, &hero.Armour, &hero.Shield); err != nil {
		return nil, err
	}

	return hero, nil
}

// Heros provides a list of heroes from the store
func (s *Store) Heros() ([]*models.Hero, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Commit()

	stmt, err := tx.Prepare(HerosCollectionQuery)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	// Scan and grow our slice of heroes. This is gnarly, but it's the best we get when using Go's sql package.
	var heros []*models.Hero
	for rows.Next() {
		var hero = &models.Hero{}
		if err := rows.Scan(&hero.ID, &hero.Name, &hero.RealName, &hero.Health, &hero.Armour, &hero.Shield); err != nil {
			return nil, err
		}

		heros = append(heros, hero)
	}

	return heros, nil
}
