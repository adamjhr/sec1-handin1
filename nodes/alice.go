package nodes

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/adamjhr/sec1handin1/calc"
	"github.com/adamjhr/sec1handin1/message"
	"github.com/adamjhr/sec1handin1/node"
)

func AliceMain(self *node.Node) {
	g := 666
	p := 6661
	bob_pk := 2227
	m := 2000
	rand.Seed(time.Now().UnixNano())
	y := rand.Intn(p)
	fmt.Println(self.Name, ": i'm encrypting the message", m, "with y =", y)

	m_encr := calc.BigModulus(bob_pk, y, p) * m
	fmt.Println(self.Name, ": the encrypted message is", m_encr)

	pk := calc.BigModulus(g, y, p)
	fmt.Println(self.Name, ": my own public key is", pk)

	message := message.Message{Sender: "Alice", Value: m_encr, Pub_key: pk}
	fmt.Println(self.Name, ": broadcasting message")

	for _, channel := range self.Outgoing {
		*channel <- message
	}
}
