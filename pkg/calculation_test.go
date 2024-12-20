package calculation_test

import (
	"testing"

	calculation "github.com/Gromov2009/calc_go/pkg"
)

func TestCalc(t *testing.T) {

	testCasesSuccess := []struct {
		name       string
		expression string
		expected   float64
	}{
		{
			name:       "simple_1plus1",
			expression: "2+3",
			expected:   5,
		},
		{
			name:       "simple_priority",
			expression: "3*4-5+1",
			expected:   8,
		},
	}

	for _, testCase := range testCasesSuccess {
		t.Run(testCase.name, func(t *testing.T) {
			val, err := calculation.Calc(testCase.expression)
			if err != nil {
				t.Fatalf("successful case %s returns error", testCase.expression)
			}
			if val != testCase.expected {
				t.Fatalf("%f should be equal %f", val, testCase.expected)
			}
		})
	}

	testCasesFail := []struct {
		name        string
		expression  string
		expectedErr error
	}{
		{
			name:       "simple",
			expression: "1+1*",
		},
		{
			name:       "priority",
			expression: "2+2**2",
		},
		{
			name:       "priority",
			expression: "((2+2-*(2",
		},
		{
			name:       "empty",
			expression: "",
		},
	}

	for _, testCase := range testCasesFail {
		t.Run(testCase.name, func(t *testing.T) {
			val, err := calculation.Calc(testCase.expression)
			if err == nil {
				t.Fatalf("expression %s is invalid but result  %f was obtained", testCase.expression, val)
			}
		})
	}

}
