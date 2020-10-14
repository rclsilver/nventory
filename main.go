package main

import (
  "github.com/rclsilver/nventory/controllers"

  "github.com/gorilla/handlers"
  "github.com/gorilla/mux"

  "fmt"
  "log"
  "net/http"
  "os"
)

func InitializeRouter() http.Handler {
  // Create the router
  router := mux.NewRouter()

  // Define API routes
  api := router.PathPrefix("/api")
  api.Methods("GET").Path("/debug/headers").Name("Debug-Headers").HandlerFunc(controllers.DebugHeaders)

  // Returns a logged router
  return handlers.LoggingHandler(os.Stdout, router)
}

func main() {
  // TODO: command-line variable or environment variable
  httpPort := 8080
  httpAddr := "0.0.0.0"

  // Configure logging
  log.SetFlags(log.Ldate | log.Ltime)

  // Initialize HTTP router
  router := InitializeRouter()

  log.Printf("Listening on %s:%d", httpAddr, httpPort)

  err := http.ListenAndServe(fmt.Sprintf("%s:%d", httpAddr, httpPort), router)

  if err != nil {
    log.Fatal(err)
  }
}
