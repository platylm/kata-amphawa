package grade_average

import "testing"

func Test_GradeAverage_InputInput_25_4_Should_Be_0(t *testing.T) {
	expected := 0.0
	inputPoint := []int{25, 4}

	actualResult := GradeAverage(inputPoint)

	if expected != actualResult {
		t.Errorf("expect %3.f but got it %3.f", expected, actualResult)
	}
}

func Test_GradeAverage_InputInput_ListPoints_Should_Be_0_Dot_7(t *testing.T) {
	expected := 1.375000
	inputPoint := Readfile("inputdata.txt")

	actualResult := GradeAverage(inputPoint)

	if expected != actualResult {
		t.Errorf("expect %f but got it %f", expected, actualResult)
	}
}
