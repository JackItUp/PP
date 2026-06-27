package main

import "fmt"

type singlyLinkedList struct {
	first *item
	last  *item
	size  int
}

type item struct {
	v    any
	next *item
}

func newSinglyLinkedList() *singlyLinkedList {
	return &singlyLinkedList{
		first: nil,
		last:  nil,
		size:  0,
	}
}

// add - добавление значения в связный список (пушбэк)
func add(l *singlyLinkedList, v any) {
	newItem := &item{v: v, next: nil}

	if l.size == 0 {
		l.first = newItem
		l.last = newItem
	} else {
		l.last.next = newItem
		l.last = newItem
	}
	l.size++
}

// addFirst - добавление значения в начало списка
func addFirst(l *singlyLinkedList, v any) {
	newItem := &item{v: v, next: l.first}
	l.first = newItem
	if l.size == 0 {
		l.last = newItem
	}
	l.size++
}

// get - получение значения по индексу из связанного списка
func get(l *singlyLinkedList, idx int) any {
	if idx < 0 || idx >= l.size {
		return nil
	}

	current := l.first
	for i := 0; i < idx; i++ {
		current = current.next
	}
	return current.v
}

// remove - удаление значения по индексу из списка
func remove(l *singlyLinkedList, idx int) bool {
	if idx < 0 || idx >= l.size {
		return false
	}

	if idx == 0 {
		l.first = l.first.next
		if l.size == 1 {
			l.last = nil
		}
		l.size--
		return true
	}

	current := l.first
	for i := 0; i < idx-1; i++ {
		current = current.next
	}

	current.next = current.next.next
	if idx == l.size-1 {
		l.last = current
	}
	l.size--
	return true
}

// values - получение слайса значений из списка
func values(l *singlyLinkedList) []any {
	result := make([]any, l.size)
	current := l.first
	for i := 0; i < l.size; i++ {
		result[i] = current.v
		current = current.next
	}
	return result
}

// size - получение размера списка
func size(l *singlyLinkedList) int {
	return l.size
}

// print - вывод списка
func printList(l *singlyLinkedList) {
	if l.size == 0 {
		fmt.Println("Список пуст")
		return
	}
	fmt.Print("Список: ")
	current := l.first
	for current != nil {
		fmt.Print(current.v, " ")
		current = current.next
	}
	fmt.Println()
}

func main() {
	fmt.Println("Список")

	l := newSinglyLinkedList()

	add(l, 100)
	add(l, 200)
	add(l, 300)
	addFirst(l, 50)

	printList(l)

	fmt.Println("Get(2):", get(l, 2))

	remove(l, 1)

	printList(l) // Список: 50 → 200 → 300 → nil
}
