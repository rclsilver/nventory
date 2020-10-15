package main

import (
  "github.com/rclsilver/nventory/authentication"
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

  // Enable external authentication
  router.Use(authentication.ExternalAuthenticationHandler)

  // Define API routes
  api := router.PathPrefix("/api").Subrouter()
  api.Methods("GET").Path("/debug/headers").HandlerFunc(controllers.DebugHeaders)
  api.Methods("GET").Path("/debug/user").HandlerFunc(controllers.DebugCurrentUser)

  // Serve frontend
  router.PathPrefix("/").Handler(http.FileServer(http.Dir("./ui")))

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
