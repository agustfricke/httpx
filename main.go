package main

import (
	"bufio"
  "mime/multipart"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/user"
	"strings"
)

func main() {
  var url string
  reader := bufio.NewReader(os.Stdin)


  for {
    if (url == "") {
      currentUser, err := user.Current()

      if err != nil {
        fmt.Println("Error:", err)
        return
      }

      fmt.Println("Hello", currentUser.Username, "welcome to snet client api!")
      fmt.Print("Enter a base URL to your session: ")
      fmt.Scanf("%v\n", &url)
    }


    var token string
    fmt.Print("Enter your token:")
    fmt.Scanf("%s", &token)

    var method string
    fmt.Print("Enter Method: POST(1) GET(2) PUT(3) DELETE(4):")
    fmt.Scanf("%v\n", &method)


    // GET DATA FOR POST AND PUT
    if method == "1" || method == "3" {

      data := make(map[string]string)

      for {
        var key string
        fmt.Print("Enter Key(leave empty to finish): ")
        key, _ = reader.ReadString('\n')
        key = strings.TrimSpace(key)

        if key == "" {
          break
        }

        var value string
        fmt.Printf("Enter Value for \"%s\": ", key)
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

        req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
        if err != nil {
          fmt.Println("Error al crear la solicitud PUT:", err)
          return
        }

        req.Header.Set("Content-Type", "application/json")
        req.Header.Set("Authorization", "Bearer "+ token)

        client := &http.Client{}
        resp, err := client.Do(req)
        if err != nil {
          fmt.Println("Error al realizar la solicitud PUT:", err)
          return
        }
        defer resp.Body.Close()

        respBody, err := ioutil.ReadAll(resp.Body)
        if err != nil {
          fmt.Println("Error al leer la respuesta:", err)
          return
        }

        fmt.Println("Respuesta del servidor:")
        fmt.Println(string(respBody))


      // PUT
      } else if method == "3" {
        var param string
        fmt.Print("Enter the ID of the record to update:")
        fmt.Scanf("%v\n", &param)
        new_url := url + param

        req, err := http.NewRequest(http.MethodPut, new_url, bytes.NewBuffer(jsonData))
        if err != nil {
          fmt.Println("Error al crear la solicitud PUT:", err)
          return
        }

        req.Header.Set("Content-Type", "application/json")
        req.Header.Set("Authorization", "Bearer "+ token)

        client := &http.Client{}
        resp, err := client.Do(req)
        if err != nil {
          fmt.Println("Error al realizar la solicitud PUT:", err)
          return
        }
        defer resp.Body.Close()

        respBody, err := ioutil.ReadAll(resp.Body)
        if err != nil {
          fmt.Println("Error al leer la respuesta:", err)
          return
        }

        fmt.Println("Respuesta del servidor:")
        fmt.Println(string(respBody))
      }


      // GET
    } else if method == "2" {
      req, err := http.NewRequest(http.MethodGet, url, nil)
      if err != nil {
        fmt.Println("Error al crear la solicitud GET:", err)
        return
      }

      req.Header.Set("Authorization", "Bearer "+ token)

      client := &http.Client{}
      resp, err := client.Do(req)
      if err != nil {
        fmt.Println("Error al realizar la solicitud GET:", err)
        return
      }
      defer resp.Body.Close()

      body, err := ioutil.ReadAll(resp.Body)
      if err != nil {
        fmt.Println("Error:", err)
        return
      }

      responseBody := string(body)

      fmt.Println(responseBody)


      // DELETE
    } else if method == "4" {
      var param string
      fmt.Print("Enter the ID of the record to update:")
      fmt.Scanf("%v\n", &param)
      new_url := url + param

      req, err := http.NewRequest(http.MethodDelete, new_url, nil)
      req.Header.Set("Authorization", "Bearer "+ token)

      if err != nil {
        fmt.Println("Error al crear la solicitud DELETE:", err)
        return
	    }

      client := &http.Client{}
      resp, err := client.Do(req)
      if err != nil {
        fmt.Println("Error al realizar la solicitud DELETE:", err)
        return
      }
      defer resp.Body.Close()

      if resp.StatusCode == http.StatusOK {
        fmt.Println("La solicitud DELETE fue exitosa.")
      } else {
        fmt.Println("La solicitud DELETE falló con el código de estado:", resp.StatusCode)
      }
    }


    var ok string
    fmt.Print("Do you want to make another request? y/n: ")
    fmt.Scanf("%v\n", &ok)

    if ok == "n" {
      fmt.Println("Bye :D")
      break
    }
  }
}


