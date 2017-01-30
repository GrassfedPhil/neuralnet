package models

import "testing"

func TestCalcNetInputReturnsNetInput(t *testing.T) {
	var node Node
	weight := []float64{2, 3}
	input := []float64{1, 2}
	bias := 1.5
	node.CalcNetInput(input, weight, bias)

	result := (weight[0] * input[0]) + (weight[1] * input[1]) + (1.5 * 1)

	if netInput := node.netInput; netInput != result {
		t.Errorf("expected output %v but got %v", result, netInput)
	}
}

func TestCalcOutputReturnsOutput(t *testing.T) {
	var node Node
	node.netInput = 1.2049573

	node.CalcOutput()

	result := 0.7694054854526742
	if output := node.output; output != result {
		t.Errorf("expected output %v but got %v", result, output)
	}
}
