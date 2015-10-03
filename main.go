package main

import (
  "./controllers"
  "github.com/zenazn/goji"
)

func main() {
    goji.Get("/ApiTop", controllers.ApiController)
    goji.Serve()
}
