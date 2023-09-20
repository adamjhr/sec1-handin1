package person

type person struct {
	name              string
	receiving_channel chan message
}
