# Heaps

- First used for heap sort
- Heap sort are good for implementing priority queues
- Heaps can be represented as complete tree(Complete tree meaning some nodes can be empty but lower right ones only)

## Max heap

- Largest key are stored on the root(on the top of the tree)

## Min heap (Opposite of max heap)

- Smallest key are stored on the root(on the top of the tree)

### Heaps are represented as array

`heap = [50, 16, 48, 14, 8, 34, 20, 9, 1, 5, 7]`

- This is possible because indices of the child can be easily calculated beased on parent's index and vice versa
- Left child is always an odd number while right child is always an even number
- For example if we want left child of an index:

`(parent index) n * 2 + 1 = x (child index)`
`(child index) x * 2 + 1 = n (parent index)`

# How to insert keys in a heap (Heapify Up) (bottom to top)

Step 1:

- Add to the bottom right of the key

Step 2:

- Rearrange a heap (bottom to top)

Time complexity of insertion heap an array is O(log n) because it depends upon height

Process of rearranging the heap is called heapify

# Steps to Extract (Heapify down) (top to bottom )

1. Take out the root
2. Fill the root index with the item of last node
3. First compare the values of left and right children
4. Compare value of largest among child and root
5. If value of one child is greater, swap larger child with parent
6. Move downwards following the same procedure

Time complexity of deletion heap an array is O(log n) because it depends upon height
