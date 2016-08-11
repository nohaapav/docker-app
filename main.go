package main

import (
    "os"
    "fmt"
    "log"
    "net/http"
)

func main() {
    port := os.Getenv("PORT")
    router := NewRouter()
    fmt.Fprintf(os.Stdout, "Listening on port [%s]\n", port)
    log.Fatal(http.ListenAndServe(":" + port, router))
}
