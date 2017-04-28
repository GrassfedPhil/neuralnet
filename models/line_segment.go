package models

//LineSegment tracks a node and it's weight to another node.
type LineSegment struct {
	ForwardNode  Node
	BackwardNode Node
	LineWeight   float64
}
