package smrsh_go

import (
	"errors"
	"fmt"
	"strconv"
)


type Calc interface {
	Calculate(operation string, args []string) (float64, error)
}

type impl struct {
}

func (i impl) Calculate(operation string, args []string)(float64, error) {
	result := 0.0
	for i, v := range args {
		value, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return 0.0, errors.New(fmt.Sprintf("invalid argument %s", v))
		}

		if i == 0 {
			result = value
			continue
		}

		switch operation {
		case "*":
			result *= value
		case "/":
			if value == 0.0 {
				return 0.0, errors.New("divide by zero")
			}
			result /= value
		case "+":
			result += value
		case "-":
			result -= value
		}
	}
	return result, nil
}

func New() Calc {
	return &impl{}
}