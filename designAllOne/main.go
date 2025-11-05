package main

type Node struct {
	prev *Node
	next *Node
	keys map[string]bool
	freq int
}

type AllOne struct {
	head    *Node
	tail    *Node
	keysMap map[string]*Node
}

func NewNode(freq int) Node {
	return Node{
		keys: make(map[string]bool),
		freq: freq,
	}
}

func Constructor() AllOne {
	head := NewNode(0)
	tail := NewNode(0)

	head.next = &tail
	tail.prev = &head

	return AllOne{
		head:    &head,
		tail:    &tail,
		keysMap: make(map[string]*Node),
	}
}

func (this *AllOne) Inc(key string) {
	if curNode, ok := this.keysMap[key]; ok {
		freq := curNode.freq
		delete(curNode.keys, key)
		nextNode := curNode.next
		if nextNode == this.tail || nextNode.freq != freq+1 {
			newNode := NewNode(freq + 1)
			newNode.keys[key] = true
			newNode.next = nextNode
			newNode.prev = curNode

			curNode.next = &newNode
			nextNode.prev = &newNode
			this.keysMap[key] = &newNode
		} else {
			nextNode.keys[key] = true
			this.keysMap[key] = nextNode
		}

		if len(curNode.keys) == 0 {
			RemoveNode(curNode)
		}
	} else {
		nextNode := this.head.next
		if nextNode == this.tail || nextNode.freq > 1 {
			newNode := NewNode(1)
			newNode.keys[key] = true
			newNode.next = nextNode
			newNode.prev = this.head
			this.head.next = &newNode
			nextNode.prev = &newNode
			this.keysMap[key] = &newNode
		} else {
			nextNode.keys[key] = true
			this.keysMap[key] = nextNode
		}
	}
}

func (this *AllOne) Dec(key string) {
	curNode, ok := this.keysMap[key]
	if !ok {
		return
	}

	delete(curNode.keys, key)
	if curNode.freq == 1 {
		delete(this.keysMap, key)
	} else {
		prevNode := curNode.prev
		if prevNode == this.head || prevNode.freq < curNode.freq-1 {
			newNode := NewNode(curNode.freq - 1)
			newNode.keys[key] = true
			newNode.next = curNode
			newNode.prev = prevNode

			prevNode.next = &newNode
			curNode.prev = &newNode
			this.keysMap[key] = &newNode
		} else {
			prevNode.keys[key] = true
			this.keysMap[key] = prevNode
		}
	}

	if len(curNode.keys) == 0 {
		RemoveNode(curNode)
	}
}

func (this *AllOne) GetMaxKey() string {
	if this.head.next == this.tail {
		return ""
	}

	for key, _ := range this.tail.prev.keys {
		return key
	}
	return ""
}

func (this *AllOne) GetMinKey() string {
	if this.tail.prev == this.head {
		return ""
	}

	for key, _ := range this.tail.keys {
		return key
	}
	return ""
}

func RemoveNode(node *Node) {
	before := node.prev
	next := node.next
	next.prev = before
	before.next = next
}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */

func main() {
	obj := Constructor()
	obj.Inc("hello")
	obj.Inc("hello")
	obj.GetMaxKey()
	obj.GetMinKey()
	obj.Inc("leet")
	obj.GetMaxKey()
	obj.GetMinKey()
}
