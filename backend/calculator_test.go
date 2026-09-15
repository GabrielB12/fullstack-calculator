package main

import "testing"

func TestAdd(t *testing.T) {
	if got := Add(2, 3); got != 5 {
		t.Errorf("Add(2,3) = %v; want 5", got)
	}
}

func TestSubtract(t *testing.T) {
	if got := Subtract(5, 3); got != 2 {
		t.Errorf("Subtract(5,3) = %v; want 2", got)
	}
}

func TestMultiply(t *testing.T) {
	if got := Multiply(4, 3); got != 12 {
		t.Errorf("Multiply(4,3) = %v; want 12", got)
	}
}

func TestDivide(t *testing.T) {
	got, err := Divide(10, 2)
	if err != nil || got != 5 {
		t.Errorf("Divide(10,2) = %v, %v; want 5, nil", got, err)
	}
}

func TestDivideByZero(t *testing.T) {
	_, err := Divide(10, 0)
	if err != ErrDivisionByZero {
		t.Errorf("expected ErrDivisionByZero, got %v", err)
	}
}

func TestPower(t *testing.T) {
	if got := Power(2, 3); got != 8 {
		t.Errorf("Power(2,3) = %v; want 8", got)
	}
}

func TestSquareRoot(t *testing.T) {
	got, err := SquareRoot(9)
	if err != nil || got != 3 {
		t.Errorf("SquareRoot(9) = %v, %v; want 3, nil", got, err)
	}
}

func TestSquareRootNegative(t *testing.T) {
	_, err := SquareRoot(-4)
	if err != ErrNegativeSqrt {
		t.Errorf("expected ErrNegativeSqrt, got %v", err)
	}
}

func TestPercentage(t *testing.T) {
	got, err := Percentage(25, 50)
	if err != nil || got != 50 {
		t.Errorf("Percentage(25,50) = %v, %v; want 50, nil", got, err)
	}
}

func TestPercentageZeroBase(t *testing.T) {
	_, err := Percentage(25, 0)
	if err != ErrZeroBase {
		t.Errorf("expected ErrZeroBase, got %v", err)
	}
}
