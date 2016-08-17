package main

import (
  "os"
  "log"
  "fmt"
  "net/http"
  "io/ioutil"
  "github.com/gorilla/mux"
)

func Info(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(os.Stdout, "I'm %s\n", Hostname())
    fmt.Fprintf(w, "I'm %s\n", Hostname())
}

func Redirect(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    service := vars["service"]
    url := Url(service, os.Getenv("PORT"))

    response, err := http.Get(url)
    if err != nil {
	log.Print(err)
    } else {
    	defer response.Body.Close()
        contents, err := ioutil.ReadAll(response.Body)
        if err != nil {
            log.Print(err)
        }       
        resp := string(contents);
        fmt.Fprintf(os.Stdout, "Redirecting from [%s] to %s \nResponse: %s", Hostname(), url, resp)
        fmt.Fprintf(w, "Redirecting from [%s] to %s \nResponse: %s", Hostname(), url, resp)
    }
}
