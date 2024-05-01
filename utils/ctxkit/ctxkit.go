package ctxkit

import (
    "math/big"
)

// Factorial calculates the factorial of a given integer n using big.Int to handle large numbers.
func Factorial(n int) *big.Int {
    result := big.NewInt(1) // Initialize result as big.Int with value 1
    if n > 0 {
        for i := 1; i <= n; i++ {
            result.Mul(result, big.NewInt(int64(i))) // result *= i
        }
    }
    return result
}
