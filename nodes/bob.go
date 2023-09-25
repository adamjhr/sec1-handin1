package nodes

import (
	"fmt"

	"github.com/adamjhr/sec1handin1/calc"
	"github.com/adamjhr/sec1handin1/node"
)

func BobMain(self *node.Node) {
	m := <-self.Incoming
	x := 66
	p := 6661

	fmt.Println(self.Name, ": received message from", m.Sender, "with pub_key", m.Pub_key, "and value", m.Value)

	fmt.Println(self.Name, ": decrypting received message")

	m_decrypted := m.Value / calc.BigModulus(m.Pub_key, x, p)

	fmt.Println(self.Name, ": decrypted message is", m_decrypted)
}
