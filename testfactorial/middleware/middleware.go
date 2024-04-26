package middleware

import (
    "context"
    "encoding/json"
    "testfactorial/types"
    "net/http"
    "github.com/julienschmidt/httprouter"
)

func ValidateInput(next httprouter.Handle) httprouter.Handle {
    return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
        var input types.Input
        err := json.NewDecoder(r.Body).Decode(&input)
        if err != nil || input.A < 0 || input.B < 0 {
            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(http.StatusBadRequest)
            json.NewEncoder(w).Encode(map[string]string{"error": "Incorrect input"})
            return
        }
        ctx := r.Context()
        ctx = context.WithValue(ctx, "input", input)
        r = r.WithContext(ctx)
        next(w, r, ps)
    }
}
