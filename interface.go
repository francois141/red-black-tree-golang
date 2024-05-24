package main

type Tree[T any] interface {
	Insert(value T)
	Delete(value T)
	Find(value T) bool
	Size() int
}

type Heap[T any] interface {
	Push(value T)
	Pop() (T, error)
	Size() int
}
