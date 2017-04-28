package main

import (
	"os"
	"strconv"
	"github.com/parlihpb/neuralnet/models"
	"fmt"
)

func main() {
	fmt.Println("Hello, world")

	networkLength, _ := strconv.Atoi(os.Args[1])
	networkHeight, _ := strconv.Atoi(os.Args[2])
	//for index := 0; networkLength < networkLength; index++ {
	//	column := models.Column{}
	//	columns[index] = column
	//	lineSegment.ForwardNode = outputNode
	//}

	previousNodes := make([]models.Node, networkHeight)
	columns := make([]models.Column, networkLength)
	//initialization of input nodes
	for index := 0; index < networkHeight; index++ {
		previousNodes[index] = models.Node{}
	}

	//for each bit of length, make a new column
	for index := 0; index < networkLength; index++ {
		column := models.Column{}


		currentNodes := make([]models.Node, networkHeight)
		for _, element := range previousNodes {
			for index := 0; index < networkHeight; index++ {
				newNode := models.Node{}
				lineSegment := models.LineSegment{ForwardNode: newNode, BackwardNode: element, LineWeight: 12}
				column.Column = append(column.Column, lineSegment)
				currentNodes[index] = newNode
			}

		}
		columns[index] = column
		previousNodes = currentNodes
	}

	//initialize output column
	outputColumn := models.Column{}
	for index := 0; index < networkLength; index++ {
		newNode := models.Node{}
		segments := columns[index].Column

		for _, element := range segments {
			segment := models.LineSegment{ForwardNode: newNode, BackwardNode: element.ForwardNode, LineWeight: 12}
			outputColumn.Column = append(outputColumn.Column, segment)
		}
	}


}
