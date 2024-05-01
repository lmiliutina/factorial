package middleware

import (
    "context"
    "encoding/json"
    "net/http"
    "github.com/julienschmidt/httprouter"
    "factorial/types"
)

func Validate(next httprouter.Handle) httprouter.Handle {
    return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
        var input types.Input
        if err := json.NewDecoder(r.Body).Decode(&input); err != nil || input.A < 0 || input.B < 0 {
            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(http.StatusBadRequest)
            json.NewEncoder(w).Encode(map[string]string{"error": "Incorrect input"})
            return
        }
        // Storing the input in the context for later retrieval
        ctx := context.WithValue(r.Context(), "input", input)
        next(w, r.WithContext(ctx), ps)
    }
}
