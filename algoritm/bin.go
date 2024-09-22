package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func lr(arr []int, n int, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && arr[left] > arr[largest] {
		largest = left
	}
	if right < n && arr[right] > arr[largest] {
		largest = right
	}
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		lr(arr, n, largest)
	}
}

func heapSort(arr []int, ascending bool) {
	n := len(arr)

	for i := n/2 - 1; i >= 0; i-- {
		lr(arr, n, i)
	}
	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		lr(arr, i, 0)
	}
	if !ascending {
		reverse(arr)
	}
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите элементы массива через пробел: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	strElements := strings.Split(input, " ")

	arr := make([]int, len(strElements))
	for i, str := range strElements {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("Ошибка: неверный ввод")
			return
		}
		arr[i] = num
	}

	fmt.Print("Выберите сортировку: 1 - по возрастанию, 2 - по убыванию: ")
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	var ascending bool
	if choice == "1" {
		ascending = true
	} else if choice == "2" {
		ascending = false
	} else {
		fmt.Println("Ошибка: неверный выбор")
		return
	}

	heapSort(arr, ascending)

	fmt.Println("Отсортированный массив:", arr)
}
