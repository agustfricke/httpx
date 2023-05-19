package req

import (
  "bytes"
  "fmt"
  "io/ioutil"
  "net/http"
  "strings"

)

func Post(url string, token string, jsonData []byte) {
  var sendFile string
  fmt.Print("Do you want to send a file? (y/n): ")
  fmt.Scanf("%s\n", &sendFile)
  sendFile = strings.ToLower(strings.TrimSpace(sendFile))

  if sendFile == "y" {
    var filePath string
    fmt.Print("Enter the file path: ")
    fmt.Scanf("%s\n", &filePath)

    File(url, jsonData, filePath)

  } else {
    req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
    if err != nil {
      fmt.Println("Error al crear la solicitud POST:", err)
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
