package models

type User struct {
  Username string `json:"user_username"`
  Email    string `json:"user_email,omitempty"`
}
