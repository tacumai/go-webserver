package main

import (
  "encoding/json"
  "net/http"
)

func main() {
  http.HandleFunc("/hoge", hoge)
  http.ListenAndServe(":9998", nil)
}

type Profile struct {
  Name    string
  Hobbies []string
}

func hoge(w http.ResponseWriter, r *http.Request) {
  profile := Profile{"Alex", []string{"snowboarding", "programming"}}
  js, err := json.Marshal(profile)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(200)
  w.Write(js)
}
