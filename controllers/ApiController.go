package controllers

import (
    "fmt"
    "net/http"
    "github.com/zenazn/goji/web"
)

func ApiController(c web.C, w http.ResponseWriter, r *http.Request) {
    fmt.Printf("%v", "top")
}
