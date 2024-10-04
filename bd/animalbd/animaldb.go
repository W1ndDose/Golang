package animalbd

import (
	"database/sql"
	"log"
)

type Animal struct {
	ID      int
	Name    string
	Species string
	Age     int
}

func AddAnimal(db *sql.DB, name, species string, age int) {
	_, err := db.Exec("INSERT INTO animals (name, species, age) VALUES ($1, $2, $3)", name, species, age)
	if err != nil {
		log.Println("Error inserting data:", err)
	}
}

func GetAllAnimals(db *sql.DB) ([]Animal, error) {
	rows, err := db.Query("SELECT id, name, species, age FROM animals")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var animals []Animal
	for rows.Next() {
		var animal Animal
		if err := rows.Scan(&animal.ID, &animal.Name, &animal.Species, &animal.Age); err != nil {
			return nil, err
		}
		animals = append(animals, animal)
	}
	return animals, nil
}

func GetSortedAnimalsByName(db *sql.DB) ([]Animal, error) {
	rows, err := db.Query("SELECT id, name, species, age FROM animals ORDER BY name ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var animals []Animal
	for rows.Next() {
		var animal Animal
		if err := rows.Scan(&animal.ID, &animal.Name, &animal.Species, &animal.Age); err != nil {
			return nil, err
		}
		animals = append(animals, animal)
	}
	return animals, nil
}

func GetSortedAnimalsByIdDESC(db *sql.DB) ([]Animal, error) {
	rows, err := db.Query("SELECT id, name, species, age FROM animals ORDER BY ID DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var animals []Animal
	for rows.Next() {
		var animal Animal
		if err := rows.Scan(&animal.ID, &animal.Name, &animal.Species, &animal.Age); err != nil {
			return nil, err
		}
		animals = append(animals, animal)
	}
	return animals, nil
}

func GetSortedAnimalsByIdASC(db *sql.DB) ([]Animal, error) {
	rows, err := db.Query("SELECT id, name, species, age FROM animals ORDER BY ID ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var animals []Animal
	for rows.Next() {
		var animal Animal
		if err := rows.Scan(&animal.ID, &animal.Name, &animal.Species, &animal.Age); err != nil {
			return nil, err
		}
		animals = append(animals, animal)
	}
	return animals, nil
}
