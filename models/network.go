package models

//Network is the input nodes, output nodes and the hidden nodes in between
type Network struct {
	Columns      []Column
	InputColumn  Column
	OutputColumn Column
}

func NewNetwork(networkHeight int, networkLength int, numInputNodes int, numOutputNodes int) Network {
	network := Network{}
	previousNodes := make([]Node, networkHeight)
	columns := make([]Column, networkLength)
	//initialization of input nodes
	for index := 0; index < networkHeight; index++ {
		previousNodes[index] = Node{}
	}
	//for each bit of length, make a new column
	for index := 0; index < networkLength; index++ {
		column := Column{}

		currentNodes := make([]Node, networkHeight)
		for _, element := range previousNodes {
			for index := 0; index < networkHeight; index++ {
				newNode := Node{}
				lineSegment := LineSegment{ForwardNode: newNode, BackwardNode: element, LineWeight: 12}
				column.Column = append(column.Column, lineSegment)
				currentNodes[index] = newNode
			}

		}
		columns[index] = column
		previousNodes = currentNodes
	}

	//initialize input column
	inputColumn := createInputColumn(numInputNodes, columns, networkLength)
	network.InputColumn = inputColumn

	//initialize output column
	outputColumn := createOutputColumn(numOutputNodes, columns, networkLength)
	network.OutputColumn = outputColumn

	return network
}
func createOutputColumn(numOutputNodes int, columns []Column, networkLength int) Column{
	outputColumn := Column{}
	for index := 0; index < numOutputNodes; index++ {
		newNode := Node{}
		segments := columns[networkLength-1].Column

		for _, element := range segments {
			segment := LineSegment{ForwardNode: newNode, BackwardNode: element.ForwardNode, LineWeight: 12}
			outputColumn.Column = append(outputColumn.Column, segment)
		}
	}

	return outputColumn
}

func createInputColumn(numInputNodes int, columns []Column, networkLength int) Column {
	inputColumn := Column{}
	for index := 0; index < numInputNodes; index++ {
		newNode := Node{}
		segments := columns[networkLength-1].Column

		for _, element := range segments {
			segment := LineSegment{ForwardNode: element.BackwardNode, BackwardNode: newNode, LineWeight: 12}
			inputColumn.Column = append(inputColumn.Column, segment)
		}
	}

	return inputColumn
}

