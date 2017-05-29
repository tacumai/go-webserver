package main

import (
  "encoding/json"
  "net/http"
)

func main() {
  http.HandleFunc("/", res)
  http.ListenAndServe(":9998", nil)
}

type Mock struct {
  Status bool `json:"status"`
  Property map[string]int `json:"property"`
}

func res(w http.ResponseWriter, r *http.Request) {
  mock := Mock{ true, map[string]int{"ver": 004, "age": 40, "sex": 1 }}
  result, err := json.Marshal(mock)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(200)
  w.Write(result)
}
