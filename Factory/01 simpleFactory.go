package _8_OOD

import "fmt"

type Operation interface {
	Operation(num1, num2 float64) float64
}

type Add struct{}

func (a *Add) Operation(num1, num2 float64) float64 {
	return num1 + num2
}

type Sub struct{}

func (s *Sub) Operation(num1, num2 float64) float64 {
	return num1 - num2
}

type Mul struct{}

func (m *Mul) Operation(num1, num2 float64) float64 {
	return num1 * num2
}

type Div struct{}

func (d *Div) Operation(num1, num2 float64) float64 {
	return num1 / num2
}

type SimpleFactoty struct{}

func (s *SimpleFactoty) createOperation(params string) Operation {
	switch params {
	case "+":
		return new(Add)
	case "-":
		return new(Sub)
	case "*":
		return new(Mul)
	case "/":
		return new(Div)
	default:
		fmt.Printf("invalid type. %s\n", params)
	}
	return nil
}
