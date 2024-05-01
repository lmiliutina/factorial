package types

import (
    //"context"
    "net/http"
    "encoding/json"
     "math/big"
)

// ExtractInput extracts the Input struct from the HTTP request.
func ExtractInput(r *http.Request) (Input, bool) {
    var input Input
    err := json.NewDecoder(r.Body).Decode(&input)
    if err != nil {
        return Input{}, false
    }
    return input, true
}

// PrepareOutput prepares the output data from calculation results.
func PrepareOutput(aFactorial, bFactorial *big.Int) Output {
    return Output{
        AFactorial: aFactorial,
        BFactorial: bFactorial,
    }
}
