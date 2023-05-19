package req

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func File(url string, jsonData []byte, filePath string) error {
  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)

  jsonPart, err := writer.CreateFormField("json")
  if err != nil {
    return err
  }
  _, err = io.Copy(jsonPart, bytes.NewReader(jsonData))
  if err != nil {
    return err
  }

  if filePath != "" {
    file, err := os.Open(filePath)
    if err != nil {
      return err
    }
    defer file.Close()

    filePart, err := writer.CreateFormFile("file", filePath)
    if err != nil {
      return err
    }

    _, err = io.Copy(filePart, file)
    if err != nil {
      return err
    }
  }

  err = writer.Close()
  if err != nil {
    return err
  }

  req, err := http.NewRequest(http.MethodPost, url, body)
  if err != nil {
    return err
  }
  defer req.Body.Close()

  req.Header.Set("Content-Type", writer.FormDataContentType())

  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    return err
  }
  defer resp.Body.Close()

  respBody, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return err
  }

  fmt.Println("Respuesta del servidor:")
  fmt.Println(string(respBody))

  return nil
}
