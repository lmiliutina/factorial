package handlers

import (
    "encoding/json"
    "testfactorial/middleware"
    "testfactorial/types"
    "testfactorial/utils"
    "net/http"
    "github.com/julienschmidt/httprouter"
)

func Calculate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    input := r.Context().Value("input").(types.Input)
    aFactorial := utils.Factorial(input.A)
    bFactorial := utils.Factorial(input.B)

    // Convert big.Int to string for JSON marshalling
    result := types.Output{
        AFactorial: aFactorial.String(),
        BFactorial: bFactorial.String(),
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(result)
}


// ValidateInput is a middleware wrapper for input validation.
func ValidateInput(next httprouter.Handle) httprouter.Handle {
    return middleware.ValidateInput(next)
}
