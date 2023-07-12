package main

import (
    "log"
    "net/http"
)

func main() {
    redirectHandler := func(w http.ResponseWriter, r *http.Request) {
        http.Redirect(w, r, "https://mico.center/" + r.RequestURI, http.StatusMovedPermanently)
    }
    http.HandleFunc("/", redirectHandler)
    err := http.ListenAndServe(":80", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

