package services

import (
    "math/big"
    "factorial/types"
    "factorial/utils/ctxkit"
)

type FactorialService struct{}

// CalculateFactorial performs factorial calculations asynchronously.
func (fs FactorialService) CalculateFactorial(input types.Input) types.Output {
    // Create channels to receive factorial results
    chA := make(chan *big.Int)
    chB := make(chan *big.Int)

    // Compute factorial of A in a goroutine
    go func() {
        chA <- ctxkit.Factorial(input.A)
    }()

    // Compute factorial of B in a goroutine
    go func() {
        chB <- ctxkit.Factorial(input.B)
    }()

    // Read the results from channels
    aFact := <-chA
    bFact := <-chB

    return types.Output{AFactorial: aFact, BFactorial: bFact}
}
