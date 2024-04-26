package middleware

import (
    "bytes"
    "testfactorial/handlers"
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/julienschmidt/httprouter"
)

func TestValidateInput(t *testing.T) {
    router := httprouter.New()
    router.POST("/test", ValidateInput(func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
        w.WriteHeader(http.StatusOK)
    }))

    input := handlers.Input{A: -1, B: 4}
    inputBytes, _ := json.Marshal(input)
    req, err := http.NewRequest("POST", "/test", bytes.NewBuffer(inputBytes))
    if err != nil {
        t.Fatal(err)
    }
    req.Header.Set("Content-Type", "application/json")

    rr := httptest.NewRecorder()
    router.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusBadRequest {
        t.Errorf("Middleware did not handle invalid input correctly: got %v want %v", status, http.StatusBadRequest)
    }
}
