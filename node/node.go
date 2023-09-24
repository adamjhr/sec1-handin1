package node

type Node struct {
	Name     string
	Incoming chan Message
	Outgoing []*chan Message
	Main     func(*Node)
}

func InitNode(name string, main func(*Node)) *Node {
	r := &Node{
		name,
		make(chan Message),
		[]Message{},
		main,
	}
	return r
}

func AddOutgoing(node *Node, toAdd *Node) *Node {
	newOutgoing := &toAdd.incoming
	append(&node.outgoing, newOutgoing)
	return node
}
