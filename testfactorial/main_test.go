package main

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/julienschmidt/httprouter"
)

// TestFactorial to ensure factorial function works correctly
func TestFactorial(t *testing.T) {
    testCases := []struct {
        name     string
        input    int
        expected int
    }{
        {"Zero", 0, 1},
        {"Positive number", 5, 120},
        {"One", 1, 1},
    }

    for _, testCase := range testCases {
        t.Run(testCase.name, func(t *testing.T) {
            result := factorial(testCase.input)
            if result != testCase.expected {
                t.Errorf("Expected factorial(%d) = %d; got %d", testCase.input, testCase.expected, result)
            }
        })
    }
}

// TestCalculateHandler to ensure the calculate endpoint works correctly
func TestCalculateHandler(t *testing.T) {
    router := httprouter.New()
    router.POST("/calculate", validateInput(calculate))

    // Create a request to pass to our handler.
    input := Input{A: 5, B: 4}
    inputBytes, _ := json.Marshal(input)
    req, err := http.NewRequest("POST", "/calculate", bytes.NewBuffer(inputBytes))
    if err != nil {
        t.Fatal(err)
    }
    req.Header.Set("Content-Type", "application/json")

    // We create a ResponseRecorder to record the response.
    rr := httptest.NewRecorder()

    // ServeHTTP directly uses the router to dispatch the request
    router.ServeHTTP(rr, req)

    // Check the status code is what we expect.
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    // Unmarshal the response body to compare as structured data
    var got, want Output
    if err := json.Unmarshal(rr.Body.Bytes(), &got); err != nil {
        t.Fatal("Could not unmarshal response:", err)
    }
    want = Output{AFactorial: 120, BFactorial: 24}
    if got != want {
        t.Errorf("Handler returned unexpected body: got %+v want %+v", got, want)
    }
}
