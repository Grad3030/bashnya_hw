package main

import "fmt"

type Stack[T any] struct { //параметр типа t
	index    int
	elements []T
}

const ALLOCATE_SIZE = 10

func (stack *Stack[T]) Push(element T) {
	stack.elements[stack.index] = element
	stack.index++
}

func (stack *Stack[T]) Pop() T {
	if stack.index > 0 {
		stack.index--
		element := stack.elements[stack.index]
		stack.elements = stack.elements[:stack.index]
		return element
	}

	panic("стек пустой")
}
func (stack *Stack[T]) IsEmpty() bool {
	if stack.index > 0 {
		return false
	} else {
		return true
	}
}
func (stack *Stack[T]) Size() int {
	return stack.index
}
func (stack *Stack[T]) Clear() {
	stack.elements = stack.elements[:0]
	stack.index = 0
}
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		elements: make([]T, ALLOCATE_SIZE),
		index:    0,
	}
}
func main() {

	stack := NewStack[int]()

	fmt.Printf("создан стек. размер: %d, проверка на пустоту: %t\n",
		stack.Size(), stack.IsEmpty())

	fmt.Println("\nдобавление элемнтов")
	for i := 1; i <= 5; i++ {
		stack.Push(i)
		fmt.Printf("добавили: %d, размер: %d\n", i, stack.Size())
	}

	fmt.Printf("\n размер=%d, проверка на пустоту=%t\n",
		stack.Size(), stack.IsEmpty())

	fmt.Println("\nизвлекаем элементы из стека")
	for !stack.IsEmpty() {
		element := stack.Pop()
		fmt.Printf("извлекли: %d, элементов в стеке: %d\n",
			element, stack.Size())
	}
	fmt.Printf("пустой ли стек %t\n", stack.IsEmpty())
}
