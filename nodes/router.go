package nodes

import (
	"fmt"

	"github.com/adamjhr/sec1handin1/node"
)

func RouterMain(self *node.Node) {
	m := <-self.Incoming
	fmt.Println(self.Name, ": received message from", m.Sender, "with pub_key", m.Pub_key, "and value", m.Value)

	fmt.Println(self.Name, ": broadcasting message")
	for _, channel := range self.Outgoing {
		*channel <- m
	}
}
