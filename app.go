package main

import (
  "fmt"
  "io/ioutil"
  "net/http"
)

func main(){
  res, err := http.Get("https://api.github.com/users/jshawl")
  contents, err := ioutil.ReadAll( res.Body )
  res.Body.Close()
  if err != nil {
    fmt.Printf( "%s", err )
  }

  fmt.Printf( "%s\n", contents )
}