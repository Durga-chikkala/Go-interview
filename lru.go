package lru

type Node struct {
	key   string
	value interface{}
	prev  *Node
	next  *Node
}

type Cache struct {
	head     *Node
	tail     *Node
	Cache    map[string]*Node
	capacity int
}

func NewCache(capacity int) *Cache {
	return &Cache{
		capacity: capacity,
		Cache:    make(map[string]*Node),
	}
}

func (c *Cache) Get(key string) (interface{}, bool) {
	if node, ok := c.Cache[key]; ok {
		c.moveToHead(node)
		return node.value, ok
	}

	return "", false
}

func (c *Cache) Put(key string, value interface{}) {
	if node, ok := c.Cache[key]; ok {
		c.Cache[key].value = value
		c.moveToHead(node)

		return
	}

	newNode := &Node{key: key, value: value}
	c.Cache[key] = newNode
	c.moveToHead(newNode)

	if c.capacity < len(c.Cache) {
		removedTailNode := c.removeTailNode()
		if removedTailNode != nil {
			delete(c.Cache, removedTailNode.key)
		}
	}
}

func (c *Cache) moveToHead(n *Node) {
	c.removeNode(n)
	c.addToHead(n)
}

func (c *Cache) removeNode(n *Node) {
	if n.prev != nil {
		n.prev.next = n.next
	} else {
		c.head = n.next
	}

	if n.next != nil {
		n.next.prev = n.prev
	} else {
		c.tail = n.prev
	}

	n.prev, n.next = nil, nil
}

func (c *Cache) addToHead(n *Node) {
	if c.head != nil {
		n.next.prev = n
		n.next = c.head
	}

	c.head = n

	if c.tail == nil {
		c.tail = n
	}
}

func (c *Cache) removeTailNode() *Node {
	if c.tail == nil {
		return nil
	}

	node := c.tail

	c.removeNode(node)

	return node
}
