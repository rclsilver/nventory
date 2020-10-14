package controllers

import (
  "net/http"
)

func DebugHeaders(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json;charset=UTF-8")
  w.WriteHeader(http.StatusOK)
  JSONEncode(w, r.Header)
}
