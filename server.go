
package main

import (
      "fmt"
      "github.com/julienschmidt/httprouter"
      "net/http"
      "log"
      "encoding/json"
)


func showSomething(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
      hello := "hello, working correctly"
      data, err := json.Marshal(hello)
      if err != nil {
          fmt.Println(err)
          return
      }
      fmt.Fprint(w, string(data))
}



func main() {

  router := httprouter.New()
  router.GET("/", showSomething)
  fmt.Println("Server up and running on port 8000")
  log.Fatal(http.ListenAndServe(":8000", router))


}
