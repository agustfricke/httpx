package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/agustfricke/snet-client-api/cli"
	"github.com/agustfricke/snet-client-api/req"

)

func main() {

  cli.Cli()
  cli.Welcome()

  var url string
  var token string
  reader := bufio.NewReader(os.Stdin)

  for {
    if (url == "") {
      fmt.Print("session url -> ")
      fmt.Scanf("%v\n", &url)
    }

    if (token == "") {
      fmt.Print("jwt token -> ")
      fmt.Scanf("%s", &token)
      token = strings.TrimRight(token, "%")
    }

    var method string
    fmt.Print("POST(1) GET(2) PUT(3) DELETE(4) -> ")
    fmt.Scanf("%v\n", &method)

    if method == "1" || method == "3" {
      data := make(map[string]string)

      for {
        var key string
        fmt.Print("key ->  ")
        key, _ = reader.ReadString('\n')
        key = strings.TrimSpace(key)

        if key == "" {
          break
        }
        var value string
        fmt.Printf("value for {\"%s\"} -> ", key)
        value, _ = reader.ReadString('\n') 
        value = strings.TrimSpace(value)
        data[key] = value
      }

      jsonData, err := json.Marshal(data)
      if err != nil {
        fmt.Println("Error converting data to JSON:", err)
        return
      }

      // POST
      if method == "1" {
        req.Post(url, token, jsonData)

        // PUT
      } else if method == "3" {
        req.Put(url, token, jsonData)
      }

      // GET
    } else if method == "2" {
      req.Get(url, token)
      // DELETE
    } else if method == "4" {
      req.Delete(url, token)
    }

    var ok string
    fmt.Print("exit? {y to exit} -> ")
    fmt.Scanf("%v\n", &ok)

    if ok == "y" {
      cli.Bye()
      break
    }

    var to string
    fmt.Print("update jwt token? {y to update} -> ")
    fmt.Scanf("%v\n", &to)

    if to == "y" {
      token = ""
    }
    
    var u string
    fmt.Print("update url {y to update} ->  ")
    fmt.Scanf("%v\n", &u)

    if(u == "y") {
      url = ""
    }

    cli.Cli()
  }
}


