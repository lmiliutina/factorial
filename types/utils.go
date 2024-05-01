package types

import (
    //"context"
    "net/http"
    //"encoding/json"
     "math/big"
)

// ExtractInput extracts the Input struct from the HTTP request.
func ExtractInput(r *http.Request) (Input, bool) {
  input, ok := r.Context().Value("input").(Input)
  return input, ok
}

// PrepareOutput prepares the output data from calculation results.
func PrepareOutput(aFactorial, bFactorial *big.Int) Output {
    return Output{
        AFactorial: aFactorial,
        BFactorial: bFactorial,
    }
}
