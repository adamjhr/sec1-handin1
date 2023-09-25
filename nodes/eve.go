package nodes

import (
	"fmt"

	"github.com/adamjhr/sec1handin1/calc"
	"github.com/adamjhr/sec1handin1/node"
)

func EveMain(self *node.Node) {
	g := 666
	p := 6661
	bob_pub_key := 2227
	private_key := 0

	fmt.Println(self.Name, ": Finding Bob's private key")

	for x := 0; x <= p; x++ {
		if calc.BigModulus(g, x, p) == bob_pub_key {
			private_key = x
			break
		}
	}

	fmt.Println(self.Name, ": Bob's private key is", private_key)

	m := <-self.Incoming

	fmt.Println(self.Name, ": intercepted message from", m.Sender, "with pub_key", m.Pub_key, "and value", m.Value)
	fmt.Println(self.Name, ": decrypting!")

	comp_key := calc.BigModulus(m.Pub_key, private_key, p)
	m_decrypted := m.Value / comp_key
	fmt.Println(self.Name, ": The decrypted message is", m_decrypted)
}
