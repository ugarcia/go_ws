/*
Go Gigs - Main Package.

Initializes database connection and http server.
 */
package main

import (
        "fmt"
    "net/http"
    // "gitlab.com/fc_freelance/go_gigs/db"
    // "gitlab.com/fc_freelance/go_gigs/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[0:])
}

// Main execution thread
func main() {

    // Init database connection
    // db.Init()

    // Init HTTP server
    // http.Init()
    
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
