package main

import "fmt"

type Employee struct {
	Name     string // имя
	Age      int    // возраст
	Position string // позиция
	Salary   int    // зарплата
}

var commands = `
1 - Добавить нового сотрудника
2 - Удалить сотрудника
3 - Вывести список сотрудников
4 - Выйти из программы
`

func main() {
	const size = 512
	empls := [size]*Employee{}

	count := 0 // счетчик сотрудников

	work := true // флаг, чтобы выйти из цикла и заверишть программу

	for work {
		cmd := 0
		fmt.Print(commands)
		fmt.Print("Введите команду: ")
		fmt.Scanf("%d", &cmd)

		switch cmd {
		case 1:
			// Добавляем нового сотрудника
			if count >= size {
				fmt.Println("База переполнена. Удалите сотрудника.")
				break
			}

			empl := new(Employee)
			fmt.Print("\nИмя: ")
			fmt.Scanf("%s", &empl.Name)
			fmt.Print("Возраст: ")
			fmt.Scanf("%d", &empl.Age)
			fmt.Print("Позиция: ")
			fmt.Scanf("%s", &empl.Position)
			fmt.Print("Зарплата: ")

			fmt.Scanf("%d", &empl.Salary)
			for i := 0; i < size; i++ {
				if empls[i] == nil {
					empls[i] = empl
					count++
					break
				}
			}

		case 2:
			fmt.Println("Удаляем сотрудника")

			var num int
			fmt.Print("Введите номер сотрудника: ")
			fmt.Scanf("%d", &num)

			if num < 1 || num > count {
				fmt.Println("Сотрудника с таким номером нет")
				break
			}

			for i := num - 1; i < count-1; i++ {
				empls[i] = empls[i+1]
			}
			empls[count-1] = nil

			// первое решение
			//for i := num - 1; i < count; i++ {
			//	// удяляем последнего сотрудника
			//	if i == count-1 {
			//		empls[i] = nil
			//		break
			//	}
			//
			//	// удаляем предпоследнего сотрудника
			//	if i == count-2 {
			//		empls[i] = empls[i+1]
			//		empls[i+1] = nil
			//		break
			//	}
			//
			//	temp := empls[i+1]
			//	empls[i] = temp
			//	empls[i+1] = empls[i+2]
			//}

			count--

		case 3:
			if empls[0] == nil {
				fmt.Println("В базе нет сотрудников")
				break
			}

			fmt.Println("Вывод сотрудников")
			fmt.Println("Всего сотрудников: ", count)

			for i := range count {
				fmt.Println("№", i+1, ". Имя: ", empls[i].Name)
			}

		case 4:
			work = !work
		}
	}
}
