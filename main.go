package main

import (
  "os"
  "fmt"
  "net/http"
  "log"
)

func getInfo(w http.ResponseWriter, r *http.Request) {
    hostname, _ := os.Hostname()
    fmt.Fprintf(os.Stdout, "I'm %s\n", hostname)
    fmt.Fprintf(w, "I'm %s\n", hostname)
}

func main() {
    port := os.Getenv("PORT")
    fmt.Fprintf(os.Stdout, "Listening on port [%s]\n", port)
    http.HandleFunc("/", getInfo)
    err := http.ListenAndServe(":" + port, nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
