package main

import (
  "log"
  "net/http"
  "github.com/julienschmidt/httprouter"
  "factorial/handlers"
  "factorial/middleware"
  "factorial/services"
)

func main() {
  router := httprouter.New()
  service := services.FactorialService{}
  handler := &handlers.Handler{
      Service: service,
  }
  router.POST("/calculate", middleware.Validate(handler.Calculate))

  log.Println("Server running on port 8989...")
  log.Fatal(http.ListenAndServe(":8989", router))
}
