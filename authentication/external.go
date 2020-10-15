package authentication

import (
  "github.com/rclsilver/nventory/models"

  "context"
  "net/http"
)

func ExternalAuthenticationHandler(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    username := r.Header.Get("X-Auth-Username")
    email := r.Header.Get("X-Auth-Email")

    if username == "" {
      w.WriteHeader(http.StatusForbidden)
    } else {
      user := models.User{Username: username, Email: email}
      ctx := context.WithValue(r.Context(), "user", user)
      r = r.WithContext(ctx)
      next.ServeHTTP(w, r)
    }
  })
}
