package models

//Node is a type of node
type Node struct {
	Output complex64
}

//CalcOutput : this method witll calculate the output based on inputs
func (n *Node) CalcOutput(input complex64, weight []complex64, bias complex64) {
	n.Output = 0
}
