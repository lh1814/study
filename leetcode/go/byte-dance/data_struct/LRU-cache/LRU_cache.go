package LRU_cache

/*
@Author: by LH
@date:  2020/7/6
@function: map + 双向链表
*/

type LRUCache struct {
	capacity int
	size     int
	value    map[int]*LinkNode
	head     *LinkNode
	tail     *LinkNode
}

type LinkNode struct {
	key, value int
	prev, next *LinkNode
}

func Constructor(capacity int) LRUCache {
	t := LRUCache{
		capacity: capacity,
		value:    map[int]*LinkNode{},
		head:     &LinkNode{},
		tail:     &LinkNode{},
	}
	t.head.next = t.tail
	t.tail.prev = t.head
	return t
}

func (this *LRUCache) Get(key int) int {

	if value, ok := this.value[key]; ok {
		this.moveToHead(value)
		return value.value
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {

	if node, ok := this.value[key]; !ok {
		node = &LinkNode{
			key:   key,
			value: value,
		}
		this.value[key] = node

		this.addToHead(node)
		this.size++
		if this.size > this.capacity {
			removed := this.removeTail()
			delete(this.value, removed.key)
			this.size--
		}

	} else {
		node.value = value
		this.moveToHead(node)
	}

	return
}

//删除链表节点
func (this *LRUCache) deleteNode(m *LinkNode) {
	m.prev.next = m.next
	m.next.prev = m.prev
}

//节点加到头节点
func (this *LRUCache) addToHead(m *LinkNode) {
	m.next = this.head.next
	this.head.next.prev = m
	m.prev = this.head
	this.head.next = m
}

//移动到头部
func (this *LRUCache) moveToHead(m *LinkNode) {
	this.deleteNode(m)
	this.addToHead(m)
}

//删除尾节点最少用
func (this *LRUCache) removeTail() *LinkNode {
	node := this.tail.prev
	this.deleteNode(node)
	return node
}
