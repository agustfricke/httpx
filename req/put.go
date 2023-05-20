package req

import (
  "bytes"
  "fmt"
  "io/ioutil"
  "net/http"
  "strings"
)

func Put(url string, token string, jsonData []byte) {

  var param string
  fmt.Print("URL: ")
  fmt.Scanf("%v\n", &param)
  new_url := url + param + "/"

  var sendFile string
  fmt.Print("Do you want to send a file? (y/n): ")
  fmt.Scanf("%s\n", &sendFile)
  sendFile = strings.ToLower(strings.TrimSpace(sendFile))

  if(sendFile == "y") {
    var filePath string
    fmt.Print("Enter the file path: ")
    fmt.Scanf("%s\n", &filePath)
    File(new_url, jsonData, filePath)
  } else {
    req, err := http.NewRequest(http.MethodPut, new_url, bytes.NewBuffer(jsonData))
    if err != nil {
      fmt.Println("Error sending the request:", err)
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
}
