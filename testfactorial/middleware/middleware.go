package middleware

import (
    "context"
    "encoding/json"
    "testfactorial/types"
    "testfactorial/utils"
    "net/http"
    "sync"
    "github.com/julienschmidt/httprouter"
)

func ValidateInput(next httprouter.Handle) httprouter.Handle {
    return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
        var input types.Input
        err := json.NewDecoder(r.Body).Decode(&input)
        if err != nil || input.A < 0 || input.B < 0 {
            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(http.StatusBadRequest)
            json.NewEncoder(w).Encode(map[string]string{"error": "Incorrect input"})
            return
        }

        // Create a context for cancellation
        ctx := r.Context()
        ctx, cancel := context.WithCancel(ctx)
        defer cancel()

        // Use a WaitGroup to wait for both goroutines to finish
        var wg sync.WaitGroup
        wg.Add(2)

        // Goroutine to calculate factorial of A
        go func() {
            defer wg.Done()
            input.AFactorial = utils.Factorial(input.A).String()
        }()

        // Goroutine to calculate factorial of B
        go func() {
            defer wg.Done()
            input.BFactorial = utils.Factorial(input.B).String()
        }()

        // Wait for both goroutines to finish
        wg.Wait()

        // Update context with input
        ctx = context.WithValue(ctx, "input", input)
        r = r.WithContext(ctx)

        // Call the next handler
        next(w, r, ps)
    }
}
