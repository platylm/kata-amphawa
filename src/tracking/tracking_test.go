package tracking

import "testing"

func Test_Tracking_Input_Expenditure_Should_Be_185_Dot_71(t *testing.T) {
	expectedExpenditure := 185.71
	input := [][]string{}
	rows1 := []string{"1", "f20", "t15", "f40", "m200", "t15"}
	rows2 := []string{"2", "f10", "t15", "f28", "g40", "f499", "t15"}
	rows3 := []string{"3"}
	rows4 := []string{"4"}
	rows5 := []string{"5", "f25", "t15", "f40", "t15"}
	rows6 := []string{"6", "t12", "f79", "g100", "f21", "t25"}
	rows7 := []string{"7", "f71"}

	input = append(input, rows1)
	input = append(input, rows2)
	input = append(input, rows3)
	input = append(input, rows4)
	input = append(input, rows5)
	input = append(input, rows6)
	input = append(input, rows7)

	actual := Tracking(input)

	if expectedExpenditure != actual {
		t.Errorf("expected %v but got it %v", expectedExpenditure, actual)
	}
}
