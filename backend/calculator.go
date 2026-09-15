package main

import (
	"errors"
	"math"
)

var ErrDivisionByZero = errors.New("division by zero")
var ErrNegativeSqrt = errors.New("cannot take square root of a negative number")
var ErrZeroBase = errors.New("percentage base cannot be zero")

func Add(a, b float64) float64 {
	return a + b
}

func Subtract(a, b float64) float64 {
	return a - b
}

func Multiply(a, b float64) float64 {
	return a * b
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivisionByZero
	}
	return a / b, nil
}

func Power(a, b float64) float64 {
	return math.Pow(a, b)
}

func SquareRoot(a float64) (float64, error) {
	if a < 0 {
		return 0, ErrNegativeSqrt
	}
	return math.Sqrt(a), nil
}

func Percentage(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrZeroBase
	}
	return (a / b) * 100, nil
}
