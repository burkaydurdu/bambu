package bambu

import "reflect"

type node[K any, V any] struct {
	key   K
	value V
	next  *node[K, V]
}

type Hash[K any, V any] struct {
	head *node[K, V]
}

func New[K any, V any]() *Hash[K, V] {
	return &Hash[K, V]{}
}

func (h *Hash[K, V]) First() (key K, value V) {
	if h.head == nil {
		return
	}

	key = h.head.key
	value = h.head.value

	return
}

func (h *Hash[K, V]) Last() (key K, value V) {
	temp := h.head

	for temp.next != nil {
		temp = temp.next
	}

	key = temp.key
	value = temp.value

	return
}

func (h *Hash[K, V]) Length() int {
	var length = 0

	temp := h.head

	if temp == nil {
		return length
	}

	length++

	for temp.next != nil {
		temp = temp.next
		length++
	}

	return length
}

func (h *Hash[K, V]) Add(key K, value V) {
	newNode := &node[K, V]{key: key, value: value}

	if h.head == nil {
		h.head = newNode
		return
	}

	temp := h.head

	for temp.next != nil {
		temp = temp.next
	}

	temp.next = newNode
}

func (h *Hash[K, V]) Remove(key K) {
	// if it is first item
	if h.head == nil || reflect.DeepEqual(h.head.key, key) {
		h.head = nil
		return
	}

	temp := h.head

	for temp.next != nil {
		if reflect.DeepEqual(temp.next.key, key) {
			temp.next = temp.next.next
			return
		}
		temp = temp.next
	}
}

func (h *Hash[K, V]) IndexOf(index int) (key K, value V) {
	if index < 0 {
		return
	}

	temp := h.head

	for index > 0 {
		temp = temp.next
		index--
	}

	if temp == nil {
		return
	}

	key = temp.key
	value = temp.value

	return
}

func (h *Hash[K, V]) Get(key K) (value V) {
	temp := h.head

	for temp != nil {
		if reflect.DeepEqual(temp.key, key) {
			value = temp.value
			return
		}
		temp = temp.next
	}

	return
}

func (h *Hash[K, V]) Clear() {
	h.head = nil
}
