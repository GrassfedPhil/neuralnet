package models

//Network is the input nodes, output nodes and the hidden nodes in between
type Network struct {
	columns    []Column
	inputNode  Node
	outputNode Node
}
