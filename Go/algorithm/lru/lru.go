package lru

import (
	"fmt"
)

type Element struct {
	prev, next *Element
	Value      int
	Key        string
}

func (e *Element) Next() *Element {
	return e.next
}

func (e *Element) Prev() *Element {
	return e.prev
}

type LRU struct {
	cache      map[string]*Element
	head, tail *Element
	capacity   int
}

// New 新建一个缓存
func New(capacity int) *LRU {
	return &LRU{make(map[string]*Element, capacity), nil, nil, capacity}
}

// Put 往缓存中放入一个键值对
func (l *LRU) Put(key string, val int) {
	if e, ok := l.cache[key]; ok {
		e.Value = val
		l.refresh(e)
		return
	}

	// 检查缓存长度
	if l.capacity == 0 {
		return
	} else if len(l.cache) >= l.capacity {
		delete(l.cache, l.tail.Key)
		l.remove(l.tail)
	}

	// 新建键值对
	e := &Element{nil, l.head, val, key}
	l.cache[key] = e
	if len(l.cache) == 1 {
		l.tail = e
	} else {
		l.head.prev = e
	}
	l.head = e
}

// refresh 刷新键值对，让用过的键值对放到链表的开头
func (l *LRU) refresh(e *Element) {
	// 如果元素已经在链表开头了就不需要刷新
	if e.prev != nil {
		e.prev.next = e.next
		if e.next != nil {
			e.next.prev = e.prev
		} else {
			l.tail = e.prev
		}
		e.prev = nil
		e.next = l.head
		l.head.prev = e
		l.head = e
	}
}

// Get 根据键名获取键值对
func (l *LRU) Get(key string) (*Element, bool) {
	if e, ok := l.cache[key]; ok {
		l.refresh(e)
		return e, true
	}
	return nil, false
}

// Delete 根据键名删除键值对
func (l *LRU) Delete(key string) bool {
	if e, ok := l.cache[key]; ok {
		l.remove(e)
		delete(l.cache, key)
		return true
	}
	return false
}

func (l *LRU) remove(e *Element) {
	if e.prev != nil {
		e.prev.next = e.next
	} else {
		l.head = e.next
	}

	if e.next != nil {
		e.next.prev = e.prev
	} else {
		l.tail = e.prev
	}
}

func (l *LRU) All() {
	for i := l.head; i != nil; i = i.Next() {
		fmt.Println(i.Key, i.Value)
	}
}