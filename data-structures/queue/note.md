# Queue
- First In First Out
## Queue Operations
- func EnQueue(data interface{}): inserts an element at the end of the queue (at rear)
- func DeQueue() (data interface{}): removes an element from the front of queue
- func Front() (data interface{}): return front element without removing it
- func Rear() (data interface{}): return back (rear) element with removing it
- func Size() int: return the number of elements in the queue
- func IsEmpty() bool: returns true if the queue is empty

## Performance
- Space complexity (for n enQueue operation) O(n)
- Time complexity of enQueue() O(1)
- Time complexity of deQueue() O(1)
- Time complexity of isEmpty() O(1)