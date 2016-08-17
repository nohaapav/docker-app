package main

import (
    "os"
    "strings"
)

func Hostname() string {
    hostname, _ := os.Hostname()
    return hostname
}

func Url(service string, port string) string {
    url := []string{"http://", service, ":", port}
    return strings.Join(url, "")    
}
