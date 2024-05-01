package handlers

import (
    "bytes"
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/julienschmidt/httprouter"
    "factorial/services"
    "factorial/types"
)

type MockCalculateService struct{}

func (m MockCalculateService) CalculateFactorial(input types.Input) types.Output {
    return types.Output{
        AFactorial: 120, // 5!
        BFactorial: 6,   // 3!
    }
}

func TestCalculateHandler(t *testing.T) {
    req, _ := http.NewRequest("POST", "/calculate", bytes.NewBufferString(`{"a": 5, "b": 3}`))
    req.Header.Set("Content-Type", "application/json")

    rr := httptest.NewRecorder()
    router := httprouter.New()
    handler := NewHandler(MockCalculateService{})
    router.POST("/calculate", handler.Calculate)

    router.ServeHTTP(rr, req)

    expected := `{"a!":120,"b!":6}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}
