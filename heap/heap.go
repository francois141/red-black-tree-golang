package heap

type Heap struct {
	arr []int
}

func NewHeap() *Heap {
	return &Heap{arr: make([]int, 0)}
}

func (heap *Heap) swap(a int, b int) {
	tmp := heap.arr[a]
	heap.arr[a] = heap.arr[b]
	heap.arr[b] = tmp
}

// TODO: Check bounds
func (heap *Heap) Push(value int) {
	heap.arr = append(heap.arr, value)

	index := len(heap.arr) - 1

	for {
		// Check
		parent := (index - 1) / 2

		// Swap if condition is violated
		if heap.arr[parent] > heap.arr[index] {
			heap.swap(index, parent)
		}

		if index == 0 {
			break
		}

		// Decrement index
		index = (index - 1) / 2
	}
}

// TODO: Check bounds
func (heap *Heap) Pop() int {
	value := heap.arr[0]

	heap.swap(0, len(heap.arr)-1)
	heap.arr = heap.arr[:len(heap.arr)-1]

	heap.heapify(0)

	return value
}

func (heap *Heap) heapify(idx int) {
	left := idx*2 + 1
	right := idx*2 + 2

	current := idx
	if left < len(heap.arr) && heap.arr[left] < heap.arr[current] {
		current = left
	}
	if right < len(heap.arr) && heap.arr[right] < heap.arr[current] {
		current = right
	}

	if current != idx {
		heap.swap(current, idx)
		heap.heapify(current)
	}
}
