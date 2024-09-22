package main

import (
	"fmt"

	"time"
)

func printNumbers() {

	for i := 1; i <= 5; i++ {

		fmt.Println(i)

		time.Sleep(time.Second)

	}

}

func timer() {

	go printNumbers()

	time.Sleep(3 * time.Second)

}
