package main

  import (
    "fmt"
    "net/http"
    "github.com/jeremydwayne/go_playground/internal/routes"
  )

  func main() {
    router := routes.NewRouter()
    port := 8080
    addr := fmt.Sprintf(":%d", port)
    fmt.Printf("Server listening on localhost%s\n", addr)
    err := http.ListenAndServe(addr, router)
    if err != nil {
      panic(err)
    }
  }
