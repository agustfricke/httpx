package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"strings"
	"time"

  "github.com/gookit/color"
)

func getTerminalSize() (int, int, error) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		return 0, 0, err
	}

	var width, height int
	fmt.Sscanf(string(out), "%d %d", &height, &width)

	return width, height, nil
}

func main() {

	// Clear the terminal
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	// Resize the terminal to full size
	resizeCmd := exec.Command("printf", "\033[8;0;0t")
	resizeCmd.Stdout = os.Stdout
	resizeCmd.Run()

	// Print content to fill the terminal
	width, height, _ := getTerminalSize()

	// Restore the terminal to its original size
	restoreCmd := exec.Command("printf", "\033[8;"+strconv.Itoa(height)+";"+strconv.Itoa(width)+"t")
	restoreCmd.Stdout = os.Stdout
	restoreCmd.Run()

  var url string
  reader := bufio.NewReader(os.Stdin)

  for {
    if (url == "") {
      currentUser, err := user.Current()

      if err != nil {
        fmt.Println("Error:", err)
        return
      }

      text := "Hello, " + currentUser.Username + "! Welcome to snet client api!\n"

      for _, char := range text {
        color.Cyan.Printf(string(char))
        time.Sleep(50 * time.Millisecond) 
      }

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
        // Verificar si se debe enviar un archivo
        var sendFile string
        fmt.Print("Do you want to send a file? (y/n): ")
        fmt.Scanf("%s\n", &sendFile)
        sendFile = strings.ToLower(strings.TrimSpace(sendFile))

        if sendFile == "y" {
          // Obtener el path del archivo
          var filePath string
          fmt.Print("Enter the file path: ")
          fmt.Scanf("%s\n", &filePath)
        // Realizar la solicitud POST
        err = sendPostRequest(url, jsonData, filePath)

        if err != nil {
          fmt.Println("Error sending POST request:", err)
          return
        }

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


// Función para enviar una solicitud POST con la opción de archivo adjunto
func sendPostRequest(url string, jsonData []byte, filePath string) error {
  // Crea un cuerpo de solicitud con formulario multipart
  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)

  // Agrega los datos JSON al formulario multipart
  jsonPart, err := writer.CreateFormField("json")
  if err != nil {
    return err
  }
  _, err = io.Copy(jsonPart, bytes.NewReader(jsonData))
  if err != nil {
    return err
  }

  // Verifica si se debe adjuntar un archivo
  if filePath != "" {
    // Abre el archivo
    file, err := os.Open(filePath)
    if err != nil {
      return err
    }
    defer file.Close()

    // Crea una parte para el archivo en el formulario multipart
    filePart, err := writer.CreateFormFile("file", filePath)
    if err != nil {
      return err
    }

    // Copia el contenido del archivo a la parte del formulario
    _, err = io.Copy(filePart, file)
    if err != nil {
      return err
    }
  }

  // Finaliza el formulario multipart
  err = writer.Close()
  if err != nil {
    return err
  }

  // Crea una solicitud POST
  req, err := http.NewRequest(http.MethodPost, url, body)
  if err != nil {
    return err
  }
  defer req.Body.Close()

  // Establece el encabezado Content-Type adecuado
  req.Header.Set("Content-Type", writer.FormDataContentType())

  // Realiza la solicitud POST
  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    return err
  }
  defer resp.Body.Close()

  // Leer la respuesta del servidor
  respBody, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return err
  }

  // Imprimir la respuesta del servidor
  fmt.Println("Respuesta del servidor:")
  fmt.Println(string(respBody))

  return nil
}
