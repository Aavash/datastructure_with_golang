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
	lastIndex := len(h.array) - 1
	left, right := getLeftChildIndex(index), getRightChildIndex(index)
	childtoCompare := 0

	// while going from top to bottom check if index of left item is not greater than size of array
	// determine which child node to compare (larger among the child should be compared)
	for left <= lastIndex {

		// when left child is only child
		// when left child is larger than right
		if left == lastIndex {
			childtoCompare = left
		} else if h.array[left] > h.array[right] {
			childtoCompare = left
		} else { // when right child is larger than left
			childtoCompare = right
		}
		// swap if the larger child has value larger than parent
		if h.array[index] < h.array[childtoCompare] {
			h.swap(index, childtoCompare)
			index = childtoCompare
			left, right = getLeftChildIndex(index), getRightChildIndex(index)
		} else {
			return
		}

	}

}

// swap keys in the array
func (h *MaxHeap) swap(index1, index2 int) {
	h.array[index1], h.array[index2] = h.array[index2], h.array[index1]
}

func Heapify() {
	heap := MaxHeap{array: []int{}}
	items := []int{1, 2, 3, 5, 7, 6}
	for _, element := range items {
		heap.Insert(element)
		fmt.Println(heap.array)
	}
	heap.Extract()
	fmt.Println(heap.array)

	heap.maxHeapifyUp(1)
}
