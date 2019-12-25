package ordereddict

type ListNode struct {
	Prev *ListNode
	Next *ListNode

	Value string
}

type DoubleLinkedList struct {
	Head *ListNode
	Tail *ListNode
}

type OrderedDict struct {
	list     *DoubleLinkedList
	listHash map[string]*ListNode

	hash map[string]string
}

func NewOrderedDict() *OrderedDict {
	return &OrderedDict{
		list:     &DoubleLinkedList{},
		listHash: make(map[string]*ListNode),
		hash:     make(map[string]string),
	}
}

func (d *OrderedDict) Get(key string) (string, bool) {
	value, ok := d.hash[key]
	return value, ok
}

func (d *OrderedDict) Add(key, value string) {
	node := &ListNode{
		Value: key,
	}

	if d.list.Head == nil {
		d.list.Head = node
	}

	if d.list.Tail == nil {
		d.list.Tail = node
	} else {
		prev := d.list.Tail
		prev.Next = node

		d.list.Tail = node
		d.list.Tail.Prev = prev
	}

	d.listHash[key] = node

	d.hash[key] = value
}

func (d *OrderedDict) Delete(key string) bool {
	node, ok := d.listHash[key]
	if !ok {
		return false
	}

	prev := node.Prev
	next := node.Next

	if prev == nil {
		//node is a head
		d.list.Head = node.Next
	} else {
		prev.Next = next
	}

	if next == nil {
		//node is a tail
		d.list.Tail = node.Prev
	} else {
		next.Prev = prev
	}

	delete(d.listHash, key)
	delete(d.hash, key)
	return true
}

func (d *OrderedDict) Keys() []string {
	out := make([]string, 0, len(d.listHash))
	node := d.list.Head
	for {
		if node == nil {
			break
		}

		out = append(out, node.Value)
		node = node.Next
	}
	return out
}