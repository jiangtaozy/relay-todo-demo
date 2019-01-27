/*
 * Maintained by jemo from 2019.1.15 to now
 * Created by jemo on 2019.1.15 10:01
 * todo server
 */

package main

import (
  "log"
  "flag"
  "net/http"
  "github.com/rs/cors"
  "github.com/jiangtaozy/relay-todo-demo/graphql"
)

var port = flag.String("port", ":3001", "server listening port")

func main() {
  graphql.Init()
  mux := http.NewServeMux()
  mux.HandleFunc("/graphql", graphql.GraphqlHandle)
  handler := cors.Default().Handler(mux)
  log.Printf("listen at %s\n", *port)
  err := http.ListenAndServe(*port, handler)
  if err != nil {
    log.Fatal("ListenAndServe error, ", err)
  }
}
