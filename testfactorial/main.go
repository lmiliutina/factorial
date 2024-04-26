package main

import (
    "context"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "github.com/julienschmidt/httprouter"
)

// Data structure to receive the input
type Input struct {
    A int `json:"a"`
    B int `json:"b"`
}

// Data structure to send the output
type Output struct {
    AFactorial int `json:"a!"`
    BFactorial int `json:"b!"`
}

// Middleware to validate input
func validateInput(next httprouter.Handle) httprouter.Handle {
    return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
        var input Input
        err := json.NewDecoder(r.Body).Decode(&input)
        if err != nil || input.A < 0 || input.B < 0 {
            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(http.StatusBadRequest)
            json.NewEncoder(w).Encode(map[string]string{"error": "Incorrect input"})
            return
        }
        ctx := context.WithValue(r.Context(), "input", input)
        next(w, r.WithContext(ctx), ps)
    }
}

// Factorial function
func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}

// Calculate handler
func calculate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    input := r.Context().Value("input").(Input)
    aFactorial := factorial(input.A)
    bFactorial := factorial(input.B)

    result := Output{AFactorial: aFactorial, BFactorial: bFactorial}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(result)
}

func main() {
    router := httprouter.New()
    router.POST("/calculate", validateInput(calculate))

    fmt.Println("Server running on port 8989...")
    log.Fatal(http.ListenAndServe(":8989", router))
}
