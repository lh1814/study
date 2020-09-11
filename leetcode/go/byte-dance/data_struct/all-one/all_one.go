package all_one

/*
@Author: by LH
@date:  2020/7/7
@function: 全O(1)数据结构
*/

type AllOne struct {
	value map[string]*Node
	size  int
	head  *Node
	tail  *Node
}

type Node struct {
	//次数
	val int
	key map[string]int
	//key num
	size int
	pre  *Node
	next *Node
}

func (n *Node) addKey(key string) {
	n.key[key] = 0
	n.size++
}

func (this *AllOne) reduceKey(n *Node, key string) (r *Node) {
	//+1
	if n.val > 1 {
		if n.pre.val != n.val-1 {
			r = &Node{
				val:  n.val - 1,
				key:  map[string]int{key: 0},
				size: 1,
			}
			this.addNodeToAimTail(r, n.pre)
		} else {
			n.pre.addKey(key)
			r = n.pre
		}
	}

	//-1
	n.size--
	delete(n.key, key)
	if n.size <= 0 {
		this.removeNode(n)
	}
	return
}

func (this *AllOne) addKey(n *Node, key string) (r *Node) {
	//+1
	if n.next.val != n.val+1 {
		r = &Node{
			val:  n.val + 1,
			key:  map[string]int{key: 0},
			size: 1,
		}
		this.addNodeToAimTail(r, n)
	} else {
		n.next.addKey(key)
		r = n.next
	}
	//-1
	n.size--
	delete(n.key, key)
	if n.size <= 0 {
		this.removeNode(n)
	}
	return
}

func (this *AllOne) removeNode(m *Node) {
	m.pre.next = m.next
	m.next.pre = m.pre
}
func (this *AllOne) addNodeToAimTail(m, aim *Node) {
	m.next = aim.next
	aim.next.pre = m
	aim.next = m
	m.pre = aim
}

/** Initialize your data structure here. */
func Constructor() AllOne {
	t := AllOne{
		value: map[string]*Node{},
		head: &Node{
			val: -1,
		},
		tail: &Node{
			val: -1,
		},
	}
	t.head.next = t.tail
	t.tail.pre = t.head
	return t
}

/** Inserts a new key <Key> with value 1. Or increments an existing key by 1. */
func (this *AllOne) Inc(key string) {
	if node, ok := this.value[key]; ok {
		this.value[key] = this.addKey(node, key)
	} else {
		this.size++
		if this.head.next.val != 1 {
			t := &Node{
				val:  1,
				key:  map[string]int{key: 0},
				size: 1,
				pre:  nil,
				next: nil,
			}
			this.value[key] = t
			this.addNodeToAimTail(t, this.head)
		} else {
			this.value[key] = this.head.next
			this.head.next.addKey(key)
		}
	}
}

/** Decrements an existing key by 1. If Key's value is 1, remove it from the data structure. */
func (this *AllOne) Dec(key string) {
	if node, ok := this.value[key]; ok {
		t := this.reduceKey(node, key)
		if t == nil {
			delete(this.value, key)
			this.size--
		} else {
			this.value[key] = t
		}
	}
}

/** Returns one of the keys with maximal value. */
func (this *AllOne) GetMaxKey() string {
	if this.size <= 0 {
		return ""
	}
	for k := range this.tail.pre.key {
		return k
	}
	return ""
}

/** Returns one of the keys with Minimal value. */
func (this *AllOne) GetMinKey() string {
	if this.size <= 0 {
		return ""
	}
	for k := range this.head.next.key {
		return k
	}
	return ""
}
