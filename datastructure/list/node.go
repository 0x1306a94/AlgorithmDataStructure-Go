package list

type _Node struct {
	element interface{}
	prev    *_Node
	next    *_Node
}

func newNode(element interface{}, prev, next *_Node) *_Node {
	return &_Node{
		element: element,
		prev:    prev,
		next:    next,
	}
}
