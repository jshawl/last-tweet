package main

import (
  "fmt"
  "os"
  "net/url"
  "github.com/ChimeraCoder/anaconda"
)

func main(){
  handle := os.Args[1]
  anaconda.SetConsumerKey("...")
  anaconda.SetConsumerSecret("...")
  api := anaconda.NewTwitterApi("...", "...")
  v := url.Values{}
  v.Set("screen_name", handle )
  res, _ := api.GetUserTimeline( v )
  tweet := res[0]
  fmt.Printf( "%s\n", tweet.Text )
}