package main

import (
	"fmt"
	"math"
)

func urav() {
	var a, b, c float64

	fmt.Print("Введите коэффициенты a, b, c: ")
	fmt.Scanln(&a, &b, &c)

	discriminant := b*b - 4*a*c

	if discriminant > 0 {
		x1 := ((-b + math.Sqrt(discriminant)) / (2 * a))
		x2 := (-b - math.Sqrt(discriminant)) / (2 * a)
		fmt.Printf("Корни уравнения: x1 = %f, x2 = %f\n", x1, x2)
	} else if discriminant == 0 {
		x := -b / (2 * a)
		fmt.Printf("Уравнение имеет один корень: x = %f\n", x)
	} else {
		fmt.Println("Уравнение не имеет действительных корней")
	}
}
