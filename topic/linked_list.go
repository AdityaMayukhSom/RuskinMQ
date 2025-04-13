package topic

type List[T any] interface {
	Length() int
	AddFront(T) T
	RemoveFront() (T, error)
	AddBack(T) T
	RemoveBack() (T, error)
	AddIndex(T, int) (T, error)
	RemoveIndex(int) (T, error)
	RemoveElem(T) (T, error)
}

type LinkedListNode[T any] struct {
	Data T
	Next *LinkedListNode[T]
}

var _ List[int] = (*LinkedList[int])(nil)

type LinkedList[T any] struct {
	head   *LinkedListNode[T]
	tail   *LinkedListNode[T]
	length int
}

// AddBack implements List.
func (ll *LinkedList[T]) AddBack(data T) T {
	node := &LinkedListNode[T]{
		Data: data,
		Next: ll.tail,
	}

	ll.tail = node
	ll.length++
	return data
}

// AddFront implements List.
func (ll *LinkedList[T]) AddFront(T) T {
	panic("unimplemented")
}

// AddIndex implements List.
func (ll *LinkedList[T]) AddIndex(T, int) (T, error) {
	panic("unimplemented")
}

// RemoveBack implements List.
func (ll *LinkedList[T]) RemoveBack() (T, error) {
	panic("unimplemented")
}

// RemoveElem implements List.
func (ll *LinkedList[T]) RemoveElem(T) (T, error) {
	panic("unimplemented")
}

// RemoveFront implements List.
func (ll *LinkedList[T]) RemoveFront() (T, error) {
	panic("unimplemented")
}

// RemoveIndex implements List.
func (ll *LinkedList[T]) RemoveIndex(int) (T, error) {
	panic("unimplemented")
}

func (ll *LinkedList[T]) Length() int {
	return ll.length
}

// func (ll *LinkedList[T]) Get() (*MessageQueue, error) {
// 	// create a new empty MessageQueue and add that to the pool and
// 	// then return the pointer to the newly created MessageQueue
// 	if mqp.head == nil {
// 		conf := &MessageQueueConfig{}
// 		newQueue, err := NewMessageQueue(conf)
// 		if err != nil {
// 			return nil, err
// 		}
// 		mqp.Add(newQueue)
// 		mqp.length++
// 	}

// 	// Get the first message queue in the pool
// 	mq := mqp.head.data

// 	if mqp.head.data.IsFull() {
// 		// Remove the first message queue from the pool when the first queue is full. This keeps on returning
// 		// the first queue unless the first queue is filled, reducing the total number of queues required.
// 		mqp.head = mqp.head.next
// 		mqp.length--
// 	}

// 	return mq, nil
// }
