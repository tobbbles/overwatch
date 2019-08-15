package sqlite

const (
	UpHeros = `CREATE TABLE IF NOT EXISTS heros  (
	id INT NOT NULL PRIMARY KEY,
	name VARCHAR(255) NOT NULL,
	real_name VARCHAR(255),
	health INT,
	armour INT,
	shield INT
	);`

	UpAbilities = `CREATE TABLE IF NOT EXISTS abilities (
	id INT NOT NULL PRIMARY KEY,
	name VARCHAR(255) NOT NULL,
	description VARCHAR(3000) NOT NULL,
	is_ultimate BOOL,
	hero INT NOT NULL,

 	FOREIGN KEY(hero) REFERENCES heros(id)
	);`
)
