package main

import (
	"errors"
	"fmt"
	"math"
)

type SqrtError struct {
	message string
}

func (s *SqrtError) Error() string {
	return s.message
}

var ErrNotFound = errors.New("Not found")

func main() {
	r := JoinedError()
	fmt.Println(r)
	if errors.Is(r, ErrNotFound) {
		fmt.Println("Unf. we dont found: ", r.Error())
	}

	var sqrtErr *SqrtError
	if errors.As(r, &sqrtErr) {
		fmt.Println("Ops has an error in math: ", sqrtErr.message)
	}

	fmt.Println()

	err := foo()
	var sqrtError *SqrtError
	if err != nil && errors.As(err, &sqrtError) {
		fmt.Println("SQRT Error", sqrtError.message)
		return
	}

	if err != nil && errors.Is(err, ErrNotFound) {
		fmt.Println("Not found error")
		return
	}

	fmt.Println("outside")
}

func foo() error {
	return &SqrtError{message: "this is not possible to do"}
}

func bar() error {
	return ErrNotFound
}

func MathSqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, &SqrtError{message: "Unable to sqrt from a number less than 0"}
	}

	result := math.Sqrt(x)

	return result, nil
}

func JoinedError() error {
	var resultErrors error

	if _, err := MathSqrt(-1); err != nil {
		resultErrors = errors.Join(resultErrors, err)
	}

	if err := bar(); err != nil {
		resultErrors = errors.Join(resultErrors, err)
	}

	return resultErrors
}
