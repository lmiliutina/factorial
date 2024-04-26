package types

// Input represents the JSON structure for incoming requests.
type Input struct {
    A int `json:"a"`
    B int `json:"b"`
}

// Output represents the JSON structure for responses.
type Output struct {
    AFactorial string  `json:"a!"`
    BFactorial string  `json:"b!"`
}
