package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method)
    if r.Method == "GET" {
        t, _ := template.ParseFiles("jsons/demo.json")
        t.Execute(w, nil)
    } else {
        r.ParseForm()
        fmt.Println("username:", r.Form["username"])
        fmt.Println("password:", r.Form["password"])
    }
}

func main() {
    http.HandleFunc("/hoge", login)

    err := http.ListenAndServeTLS(":13502", "ssl/development/myself.crt", "ssl/development/myself.key", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
  }
