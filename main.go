package main

import "fmt"

type queue struct {
	s         []any
	low, high int
	size      int
}

func newQueue(size int) *queue {
	return &queue{
		s:    make([]any, size),
		size: size,
		low:  -1,
		high: -1,
	}
}

// push - добавление в очередь значения
func push(q *queue, v any) bool {
	if q.high == q.size-1 {
		return false // очередь переполнена
	}
	if q.low == -1 {
		q.low = 0
	}
	q.high++
	q.s[q.high] = v
	return true
}

// pop - получения значения из очереди и его удаление
func pop(q *queue) any {
	if q.low == -1 || q.low > q.high {
		return nil // очередь пуста
	}
	v := q.s[q.low]
	q.low++
	return v
}

// peek - просмотр первого элемента очереди
func peek(q *queue) any {
	if q.low == -1 || q.low > q.high {
		return nil
	}
	return q.s[q.low]
}

// isEmpty - проверка очереди на пустоту
func isEmpty(q *queue) bool {
	return q.low == -1 || q.low > q.high
}

// isFull - проверка очереди на заполненность
func isFull(q *queue) bool {
	return q.high == q.size-1
}

// printQueue - вывод очереди
func printQueue(q *queue) {
	if q.low == -1 || q.low > q.high {
		fmt.Println("Очередь пуста")
		return
	}
	fmt.Print("Очередь (слева направо): ")
	for i := q.low; i <= q.high; i++ {
		fmt.Print(q.s[i], " ")
	}
	fmt.Println()
}
func main() {
	fmt.Println("Очередь")

	q := newQueue(5)

	push(q, "a")
	push(q, "b")
	push(q, "c")

	printQueue(q)

	fmt.Println("Pop:", pop(q))
	printQueue(q)
}
