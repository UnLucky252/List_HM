package list

import "fmt"

type List struct {
	len       int64
	firstNode *node
}

// NewList создает новый список
func NewList() (l *List) {
	return &List{}
}

// Len возвращает длину списка
func (l *List) Len() (len int64) {
	return l.len
}

// Add добавляет элемент в список и возвращает его индекс
func (l *List) Add(value int64) (index int64) {
	newNode := &node{l.len, value, nil}
	if l.len > 0 {
		currentNode := l.firstNode
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = newNode
	} else {
		l.firstNode = newNode
	}
	l.len++
	return newNode.index
}

// RemoveByIndex удаляет элемент из списка по индексу
func (l *List) RemoveByIndex(id int64) {
	if id < 0 || id >= l.len {
		return
	}

	if id == 0 {
		l.firstNode = l.firstNode.next
	} else {
		previousNode := l.firstNode
		for i := int64(0); i < id-1; i++ {
			previousNode = previousNode.next
		}
		previousNode.next = previousNode.next.next
	}

	l.len--
}

// RemoveByValue удаляет элемент из списка по значению
func (l *List) RemoveByValue(value int64) {
	previousNode := l.firstNode
	currentNode := l.firstNode

	for currentNode != nil {
		if currentNode.value == value {
			if currentNode == l.firstNode {
				l.firstNode = currentNode.next
			} else {
				previousNode.next = currentNode.next
			}
			l.len--
			return
		}

		previousNode = currentNode
		currentNode = currentNode.next
	}
}

// RemoveAllByValue удаляет все элементы из списка по значению
func (l *List) RemoveAllByValue(value int64) {
	currentNode := l.firstNode
	previousNode := &node{next: currentNode}

	for currentNode != nil {
		if currentNode.value == value {
			if currentNode == l.firstNode {
				l.firstNode = l.firstNode.next
				l.len--
				currentNode = l.firstNode
				previousNode = &node{next: currentNode}
			} else {
				previousNode.next = currentNode.next
				l.len--
				currentNode = currentNode.next
			}
		} else {
			previousNode = currentNode
			currentNode = currentNode.next
		}
	}
	if l.len == 0 {
		return
	}
}

// GetByIndex возвращает значение элемента по индексу.
//
// Если элемента с таким индексом нет, то возвращается 0 и false.
func (l *List) GetByIndex(id int64) (value int64, ok bool) {
	if id < 0 || id >= l.len {
		return 0, false
	}

	currentNode := l.firstNode
	for i := int64(0); i < id; i++ {
		currentNode = currentNode.next
	}

	return currentNode.value, true
}

// GetByValue возвращает индекс первого найденного элемента по значению.
//
// Если элемента с таким значением нет, то возвращается 0 и false.
func (l *List) GetByValue(value int64) (id int64, ok bool) {
	currentNode := l.firstNode
	index := int64(0)

	for currentNode != nil {
		if currentNode.value == value {
			return index, true
		}

		currentNode = currentNode.next
		index++
	}

	return 0, false
}

// GetAllByValue возвращает индексы всех найденных элементов по значению
//
// Если элементов с таким значением нет, то возвращается nil и false.
func (l *List) GetAllByValue(value int64) (ids []int64, ok bool) {
	if l.len == 0 {
		return nil, false
	}

	var indexes []int64
	currentNode := l.firstNode
	index := int64(0)

	for currentNode != nil {
		if currentNode.value == value {
			indexes = append(indexes, index)
		}

		currentNode = currentNode.next
		index++
	}

	if len(indexes) == 0 {
		return nil, false
	}

	return indexes, true
}

// GetAll возвращает все элементы списка
//
// Если список пуст, то возвращается nil и false.
func (l *List) GetAll() (values []int64, ok bool) {
	if l.len == 0 {
		return nil, false
	}

	var newValues []int64
	currentNode := l.firstNode

	for currentNode != nil {
		newValues = append(newValues, currentNode.value)
		currentNode = currentNode.next
	}

	return newValues, true
}

// Clear очищает список
func (l *List) Clear() {
	l.firstNode = nil
	l.len = 0
}

// Print выводит список в консоль
func (l *List) Print() {
	for l.firstNode != nil {
		fmt.Println(l.firstNode.value)
		l.firstNode = l.firstNode.next
	}
}
