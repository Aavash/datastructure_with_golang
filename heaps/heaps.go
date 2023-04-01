package heaps

import "fmt"

func getParentIndex(index int) int {
	return (index - 1) / 2
}

func getLeftChildIndex(index int) int {
	return 2*index + 1
}

func getRightChildIndex(index int) int {
	return 2*index + 2
}

// Maxheap struct has a slice comprising of the heap
type MaxHeap struct {
	array []int
}

// Insert method adds element to he heap (bottom-top is done)
func (h *MaxHeap) Insert(element int) {
	h.array = append(h.array, element)
	h.maxHeapifyUp(len(h.array) - 1)
}

// Extract removes the largest key from heap and heapifies
func (h *MaxHeap) Extract() int {
	extracted := h.array[0]

	length := len(h.array) - 1

	if length == 0 {
		fmt.Println("cannot extract because array length is 0")
		return -1
	}
	// take out last index and put it in root
	h.array[0] = h.array[length]
	h.array = h.array[:length]

	h.maxHeapifyDown(0)

	return extracted
}

// It will heapify from bottom top
func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[getParentIndex(index)] < h.array[index] {
		h.swap(getParentIndex(index), index)
		index = getParentIndex(index)
	}
}

// It will heapify from top to bottom
func (h *MaxHeap) maxHeapifyDown(index int) {
	for h.array[getParentIndex(index)] < h.array[index] {
		h.swap(getParentIndex(index), index)
		index = getParentIndex(index)
	}
}

// swap keys in the array
func (h *MaxHeap) swap(index1, index2 int) {
	h.array[index1], h.array[index2] = h.array[index2], h.array[index1]
}

func Heapify() {
	heap := MaxHeap{array: []int{}}
	items := []int{1, 2, 3, 5, 7, 3, 6}
	for _, element := range items {
		heap.Insert(element)
		fmt.Println(heap.array)
	}
	heap.maxHeapifyUp(1)
}
