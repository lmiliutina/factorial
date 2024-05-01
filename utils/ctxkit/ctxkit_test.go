package ctxkit

import (
    "math/big"
    "testing"
)

func TestFactorial(t *testing.T) {
    testCases := []struct {
        name     string
        input    int
        expected *big.Int
    }{
        {"0!", 0, big.NewInt(1)},
        {"1!", 1, big.NewInt(1)},
        {"5!", 5, big.NewInt(120)},
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            if got := Factorial(tc.input); got.Cmp(tc.expected) != 0 {
                t.Errorf("Factorial(%d) = %v, want %v", tc.input, got, tc.expected)
            }
        })
    }
}
