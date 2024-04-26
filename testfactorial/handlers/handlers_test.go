package handlers

import (
    "bytes"
    "encoding/json"
    "testfactorial/utils"
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/julienschmidt/httprouter"
)

func TestCalculateHandler(t *testing.T) {
    router := httprouter.New()
    router.POST("/calculate", ValidateInput(Calculate))

    input := Input{A: 5, B: 4}
    inputBytes, _ := json.Marshal(input)
    req, err := http.NewRequest("POST", "/calculate", bytes.NewBuffer(inputBytes))
    if err != nil {
        t.Fatal(err)
    }
    req.Header.Set("Content-Type", "application/json")

    rr := httptest.NewRecorder()
    router.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    var output Output
    if err := json.Unmarshal(rr.Body.Bytes(), &output); err != nil {
        t.Fatal("Could not parse response:", err)
    }

    expected := Output{AFactorial: utils.Factorial(5), BFactorial: utils.Factorial(4)}
    if output != expected {
        t.Errorf("Handler returned unexpected body: got %+v want %+v", output, expected)
    }
}
