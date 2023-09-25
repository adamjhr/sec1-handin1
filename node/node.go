package node

import "github.com/adamjhr/sec1handin1/message"

type Node struct {
	Name     string
	Incoming chan message.Message
	Outgoing []*chan message.Message
}

func InitNode(name string) *Node {
	r := &Node{
		name,
		make(chan message.Message),
		[]*chan message.Message{},
	}
	return r
}

func AddOutgoing(node *Node, toAdd *Node) {
	newOutgoing := &toAdd.Incoming
	outgoingList := &node.Outgoing
	newOutgoingList := append(*outgoingList, newOutgoing)
	node.Outgoing = newOutgoingList
}
