package main

import (
  "log"
  "net/http"
  "fmt"
)

const addr = ":8080"

func main() {
  http.HandleFunc("/", serveHome)
  http.HandleFunc("/ws", serveWs)
  if err := http.ListenAndServe(addr, nil); err != nil {
    log.Fatal(err)
  }
}

// Serves static empty-ish home
func serveHome(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path != "/" {
    http.Error(w, "Not found", 404)
    return
  }
  if r.Method != "GET" {
    http.Error(w, "Method not allowed", 405)
    return
  }
  w.Header().Set("Content-Type", "text/html; charset=utf-8")
  
  fmt.Fprintf(w, homeHTML)
}

const homeHTML = `<!DOCTYPE html>
<html lang="en">
    <head>
        <title>Empty page</title>
    </head>
    <body>
        There is nothing here for you!
    </body>
</html>
`