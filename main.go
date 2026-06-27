package main

import "fmt"

type stack struct {
	s    []any // слайс в котором хранятся значения в стеке
	head int   // индекс головы стека
}

func newStack(size int) *stack {
	return &stack{
		s:    make([]any, size),
		head: -1,
	}
}

// push - добавление в стек значения
func push(s *stack, v any) bool {
	if s.head == len(s.s)-1 {
		return false // стек переполнен
	}
	s.head++
	s.s[s.head] = v
	return true
}

// pop - получения значения из стека и его удаление из вершины
func pop(s *stack) any {
	if s.head == -1 {
		return nil // стек пуст
	}
	v := s.s[s.head]
	s.head--
	return v
}

// peek - просмотр значения на вершине стека
func peek(s *stack) any {
	if s.head == -1 {
		return nil
	}
	return s.s[s.head]
}

// isEmpty - проверка стека на пустоту
func isEmpty(s *stack) bool {
	return s.head == -1
}

// isFull - проверка стека на заполненность
func isFull(s *stack) bool {
	return s.head == len(s.s)-1
}

// printStack - вывод стека
func printStack(s *stack) {
	if s.head == -1 {
		fmt.Println("Стек пуст")
		return
	}

	fmt.Print("Стек (сверху вниз): ")
	for i := s.head; i >= 0; i-- {
		fmt.Print(s.s[i], " ")
	}
	fmt.Println()
}

func main() {
	// Стек
	fmt.Println("Стек")

	s := newStack(5)

	fmt.Println("isEmpty:", isEmpty(s))
	fmt.Println("isFull:", isFull(s))

	push(s, 10)
	push(s, 20)
	push(s, 30)

	printStack(s)

	fmt.Println("Pop:", pop(s))
	fmt.Println("Peek:", peek(s))

	printStack(s)
}
