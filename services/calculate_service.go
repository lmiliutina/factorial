package services

import "factorial/types"

// CalculateService defines the interface for calculation services.
type CalculateService interface {
    CalculateFactorial(input types.Input) types.Output
}
