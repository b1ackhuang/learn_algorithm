package _map

type LRUCache struct {
	cap                  int
	mapIndex             map[int]*DLinkedNode
	dummyHead, dummyTail *DLinkedNode
}

type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

func Constructor(capacity int) LRUCache {
	var head, tail = &DLinkedNode{key: -1, value: -1}, &DLinkedNode{key: -1, value: -1}
	head.next = tail
	tail.prev = head
	return LRUCache{
		cap:       capacity,
		mapIndex:  map[int]*DLinkedNode{},
		dummyHead: head,
		dummyTail: tail,
	}
}

func (this *LRUCache) Get(key int) int {
	tmpNode, ok := this.mapIndex[key]
	if !ok {
		return -1
	}
	this.move2Head(tmpNode)
	return tmpNode.value
}

func (this *LRUCache) Put(key int, value int) {
	tmpNode, ok := this.mapIndex[key]
	if ok {
		tmpNode.value = value
		this.move2Head(tmpNode)
		return
	}

	var newNode = &DLinkedNode{key: key, value: value}
	this.mapIndex[key] = newNode
	this.insert2Head(newNode)
	if len(this.mapIndex) > this.cap {
		var delNode = this.dummyTail.prev
		delete(this.mapIndex, delNode.key)
		this.unlinkNode(delNode)
	}
}

func (this *LRUCache) move2Head(node *DLinkedNode) {
	//unlink
	this.unlinkNode(node)
	//insert to head
	this.insert2Head(node)
}

func (this *LRUCache) insert2Head(node *DLinkedNode) {
	//insert to head
	// next ptr
	var readHead = this.dummyHead.next
	this.dummyHead.next = node
	node.next = readHead
	//pre ptr
	readHead.prev = node
	node.prev = this.dummyHead
}

func (this LRUCache) unlinkNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
	node.next = nil
	node.prev = nil
}
