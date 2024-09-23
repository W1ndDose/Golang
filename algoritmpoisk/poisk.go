package main

import (
	"container/heap"
	"fmt"
)

type Task struct {
	name     string
	priority int
}

type PriorityQueue []*Task

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {

	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Task))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	task := old[n-1]
	*pq = old[0 : n-1]
	return task
}

func main() {
	tasks := []Task{
		{"реализация интерфейса", 3},
		{"реализация функции аутентификации", 1},
		{"реализация запросов к БД", 3},
		{"подключение к БД", 2},
		{"тестирование программы", 5},
		{"реализация основной функции", 0},
		{"ревью кода", 4},
	}

	pq := make(PriorityQueue, 0)
	for _, task := range tasks {
		heap.Push(&pq, &task)
	}

	var searchName string
	fmt.Println("Введите название задачи для поиска:")
	fmt.Scanln(&searchName)

	for _, task := range pq {
		if task.name == searchName {
			fmt.Printf("Приоритет задачи '%s': %d\n", task.name, task.priority)
			break
		}
	}

	var searchPriority int
	fmt.Println("Введите номер приоритета для поиска:")
	fmt.Scanln(&searchPriority)

	for _, task := range pq {
		if task.priority == searchPriority {
			fmt.Printf("Задача с приоритетом %d: %s\n", searchPriority, task.name)
			break
		}
	}
}
