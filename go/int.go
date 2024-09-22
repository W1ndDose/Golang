package main

import (
	"fmt"
)

type Animal interface {
	Sound() string
	Move() string
	Eat() string
	Age() int
}

type Lion struct {
	age int
}

func (l *Lion) Sound() string {
	return "Рык"
}

func (l *Lion) Move() string {
	return "Бежит"
}

func (l *Lion) Eat() string {
	return "Ест мясо"
}

func (l *Lion) Age() int {
	return l.age
}

type Elephant struct {
	age int
}

func (e *Elephant) Sound() string {
	return "Трубит"
}

func (e *Elephant) Move() string {
	return "Идет медленно"
}

func (e *Elephant) Eat() string {
	return "Ест траву"
}

func (e *Elephant) Age() int {
	return e.age
}

type Bird struct {
	age int
}

func (b *Bird) Sound() string {
	return "Чирик"
}

func (b *Bird) Move() string {
	return "Летит"
}

func (b *Bird) Eat() string {
	return "Ест семена"
}

func (b *Bird) Age() int {
	return b.age
}

type Fish struct {
	age int
}

func (f *Fish) Sound() string {
	return "Булькает"
}

func (f *Fish) Move() string {
	return "Плавает"
}

func (f *Fish) Eat() string {
	return "Ест водоросли"
}

func (f *Fish) Age() int {
	return f.age
}

type Dog struct {
	age int
}

func (d *Dog) Sound() string {
	return "Гавкает"
}

func (d *Dog) Move() string {
	return "Плавает быстро"
}

func (d *Dog) Eat() string {
	return "Ест рыбу"
}

func (d *Dog) Age() int {
	return d.age
}

func inputAnimal() Animal {
	var animalType string
	var age int

	fmt.Print("Введите тип животного (Лев, Слон, Птица, Рыба, Собака): ")
	fmt.Scan(&animalType)
	fmt.Print("Введите возраст животного: ")
	fmt.Scan(&age)

	switch animalType {
	case "Лев":
		return &Lion{age: age}
	case "Слон":
		return &Elephant{age: age}
	case "Птица":
		return &Bird{age: age}
	case "Рыба":
		return &Fish{age: age}
	case "Собака":
		return &Dog{age: age}
	default:
		fmt.Println("Неизвестный тип животного!")
		return nil
	}
}

func main() {
	var animals []Animal
	var count int

	fmt.Print("Сколько животных вы хотите добавить? ")
	fmt.Scan(&count)

	for i := 0; i < count; i++ {
		fmt.Printf("Введите данные для животного %d:\n", i+1)
		animal := inputAnimal()
		if animal != nil {
			animals = append(animals, animal)
		}
	}

	fmt.Println("\nИнформация о животных:")
	for _, animal := range animals {
		fmt.Printf("Животное: %T\n", animal)
		fmt.Println("Звук:", animal.Sound())
		fmt.Println("Движение:", animal.Move())
		fmt.Println("Еда:", animal.Eat())
		fmt.Printf("Возраст: %d\n", animal.Age())
		fmt.Println()
	}
}
