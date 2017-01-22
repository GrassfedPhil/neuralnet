package models

import "testing"

func TestOutputReturnsOutput(t *testing.T) {
	var node Node
	weight := []complex64{1.4999}
	node.CalcOutput(1.5, weight, 1.5)

	if output := node.Output; output != 20 {
		t.Errorf("expected output stuff but got %v", output)
	}
}
