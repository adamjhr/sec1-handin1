package node

type Node struct {
	Name     string
	Incoming chan Message
	Outgoing []*chan Message
	Main     func(*Node)
}

func initNode(name string, main func(*Node)) *Node {
	r := &Node{
		name,
		make(chan Message),
		[]Message{},
		main,
	}
	return r
}

func addOutgoing(node *Node, toAdd *Node) *Node {
	newOutgoing := &toAdd.incoming
	append(&node.outgoing, newOutgoing)
	return node
}
