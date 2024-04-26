package main

import (
    "fmt"
    "log"
    "net/http"
    "testfactorial/handlers"
    "github.com/julienschmidt/httprouter"
)

func main() {
    router := httprouter.New()
    router.POST("/calculate", handlers.ValidateInput(handlers.Calculate))

    fmt.Println("Server running on port 8989...")
    log.Fatal(http.ListenAndServe(":8989", router))
}
