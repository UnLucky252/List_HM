package main

import (
	"fmt"
	"list"
)

func main() {
	l := list.NewList()

	l.Add(7)
	l.Add(2)
	l.Add(1)
	l.Add(5)
	l.Add(11)
	l.Add(27)
	l.Add(11)
	l.Add(3)

	fmt.Println("Длина листа:", l.Len())

	values, ok := l.GetAll()
	if ok {
		fmt.Println("Элементы листа:", values)
	} else {
		fmt.Println("Пусто")
	}

	l.RemoveByIndex(2)
	fmt.Println("Длина листа после удаления элемента по индексу:", l.Len())

	values, ok = l.GetAll()
	if ok {
		fmt.Println("Элементы в листе после удаления:", values)
	} else {
		fmt.Println("Пусто")
	}
	l.RemoveAllByValue(11)
	l.RemoveByValue(5)
	values, ok = l.GetAll()
	if ok {
		fmt.Println("Элементы в листе после удаления по значению:", values)
	} else {
		fmt.Println("Пусто")
	}

	fmt.Println(l.GetByIndex(3))    /// Получаем значение по индексу
	fmt.Println(l.GetByValue(27))   /// Получаем индекс по значению
	fmt.Println(l.GetAllByValue(2)) /// Получаем индексы всех элементов по значению

	fmt.Println("Вывод элементов листа:")
	l.Print()

	l.Clear()
	fmt.Println("Длина листа после очистки:", l.Len())
}
