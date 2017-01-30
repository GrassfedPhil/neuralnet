package models

import "math"

//Node is a type of node
type Node struct {
	netInput float64
	output   float64
}

//CalcNetInput : this method will calculate the netInput based on inputs
func (n *Node) CalcNetInput(input []float64, weight []float64, bias float64) {
	var total float64
	for i := 0; i < len(input); i++ {
		total += input[i] * weight[i]
	}
	total += (bias * 1)
	n.netInput = total
}

//CalcOutput : this method will calculate the output based on the nodes netInput
func (n *Node) CalcOutput() {
	n.output = 1 / (1 + math.Pow(math.E, -n.netInput))
}
