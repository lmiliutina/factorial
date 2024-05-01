package handlers

import (
    "encoding/json"
    "net/http"
    "github.com/julienschmidt/httprouter"
    "factorial/services"
    "factorial/types"
)

// Assuming you have a service instance passed in some way, possibly through a struct
type Handler struct {
    Service services.CalculateService
}

// Calculate handles the calculation endpoint.
func (h *Handler) Calculate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    input, ok := r.Context().Value("input").(types.Input)
    if !ok {
        http.Error(w, "Invalid request context", http.StatusBadRequest)
        return
    }

    output := h.Service.CalculateFactorial(input)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(output)
}
