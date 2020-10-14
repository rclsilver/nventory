package controllers

import (
  "encoding/json"
  "net/http"
)

func JSONEncode(writer http.ResponseWriter, value interface{}) error {
  encoder := json.NewEncoder(writer)
  encoder.SetIndent("", "  ")
  return encoder.Encode(value)
}
