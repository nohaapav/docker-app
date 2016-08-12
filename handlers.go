package main

import (
  "os"
  "log"
  "fmt"
  "strings"
  "net/http"
  "io/ioutil"
  "github.com/gorilla/mux"
)

func Info(w http.ResponseWriter, r *http.Request) {
    hostname, _ := os.Hostname()
    fmt.Fprintf(os.Stdout, "I'm %s\n", hostname)
    fmt.Fprintf(w, "I'm %s\n", hostname)
}

func Redirect(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    service := vars["service"]
    url := []string{"http://", service, ":", os.Getenv("PORT")}

    response, err := http.Get(strings.Join(url, ""))
    if err != nil {
	log.Print(err)
    } else {
    	defer response.Body.Close()
        contents, err := ioutil.ReadAll(response.Body)
        if err != nil {
            log.Print(err)
        }        
        fmt.Fprintf(os.Stdout, "Redirect to [%s] service. Response: %s\n", service, string(contents))
        fmt.Fprintf(w, "Redirect to [%s] service. Response: %s\n", service, string(contents))
    }
}
