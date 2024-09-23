package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

type Animal interface {
	Sound() (string, error)
}

type Dog struct{}

func (d Dog) Sound() (string, error) {
	return "Гав!", nil
}

type Cat struct{}

func (c Cat) Sound() (string, error) {
	return "Мяу!", nil
}

type UnknownAnimal struct{}

func (u UnknownAnimal) Sound() (string, error) {
	return "", errors.New("Неизвестный звук животного")
}

func getAnimalType() (Animal, error) {
	var animalType string
	fmt.Print("Введите тип животного (dog/cat/unknown): ")
	_, err := fmt.Scan(&animalType)
	if err != nil {
		return nil, err
	}

	switch animalType {
	case "dog":
		return Dog{}, nil
	case "cat":
		return Cat{}, nil
	case "unknown":
		return UnknownAnimal{}, nil
	default:
		return nil, errors.New("недопустимый тип животного")
	}
}

func logError(err error) {
	f, fileErr := os.OpenFile("errors.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if fileErr != nil {
		fmt.Println("Ошибка при открытии файла лога:", fileErr)
		return
	}
	defer f.Close()

	logger := log.New(f, "ERROR: ", log.LstdFlags)
	logger.Println(err)
}

func main() {
	animal, err := getAnimalType()
	if err != nil {
		logError(err)
		fmt.Println("Ошибка:", err)
		return
	}

	sound, err := animal.Sound()
	if err != nil {
		logError(err)
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Println("Звук животного:", sound)
}
