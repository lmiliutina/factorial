package middleware

import (
    "bytes"
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/julienschmidt/httprouter"
    "factorial/types"
)

func TestValidateMiddleware(t *testing.T) {
    // Test data
    tests := []struct {
        payload string
        status  int
    }{
        {`{"a": 5, "b": 3}`, http.StatusOK},
        {`{"a": -1, "b": 3}`, http.StatusBadRequest},
        {`{"a": 5, "b": -3}`, http.StatusBadRequest},
        {`{"a": "5", "b": 3}`, http.StatusBadRequest},
    }

    for _, test := range tests {
        req, _ := http.NewRequest("POST", "/calculate", bytes.NewBufferString(test.payload))
        req.Header.Set("Content-Type", "application/json")
        rr := httptest.NewRecorder()
        router := httprouter.New()
        router.POST("/calculate", Validate(func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
            w.WriteHeader(http.StatusOK)
        }))

        router.ServeHTTP(rr, req)

        if status := rr.Code; status != test.status {
            t.Errorf("handler returned wrong status code: got %v want %v",
                status, test.status)
        }
    }
}
