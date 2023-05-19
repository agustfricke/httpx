package main

import (
  "bytes"
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
  "os/user"
)

func main() {

  var url string

  for {
    // Select URL
    if (url == "") {
    currentUser, err := user.Current()
    if err != nil {
      fmt.Println("Error:", err)
      return
    }

    fmt.Println("Hello", currentUser.Username, "welcome to snet client api!")

    if (url == "") {
      fmt.Print("Enter a base URL to your session: ")
      fmt.Scanf("%v\n", &url)
    }
    }


    var token string
    fmt.Print("Enter your token:")
    fmt.Scanf("%s", &token)

    var method string
    fmt.Print("Enter Method: POST(1) GET(2) PUT(3) DELETE(4):")
    fmt.Scanf("%v\n", &method)

    if method == "1" || method == "3" {
      // Create an empty map to store the data
      data := make(map[string]string)

      for {
        // Enter data key
        var key string
        fmt.Print("Enter Key (leave empty to finish): ")
        fmt.Scanf("%s", &key)

        // Check if the key is empty
        if key == "" {
          break
        }

        // Enter value for the key
        var value string
        fmt.Printf("Enter Value for \"%s\": ", key)
        fmt.Scanf("%s", &value)

        // Add key-value pair to the data map
        data[key] = value
      }

      // Convert the data to JSON
      jsonData, err := json.Marshal(data)
      if err != nil {
        fmt.Println("Error converting data to JSON:", err)
        return
      }

      // Send the POST request
      if method == "1" {
      resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
      if err != nil {
        fmt.Println("Error sending POST request:", err)
        return
      }
      defer resp.Body.Close()

      // Read the response from the server
      respBody, err := ioutil.ReadAll(resp.Body)
      if err != nil {
        fmt.Println("Error reading response:", err)
        return
      }
      // Print the server response
      fmt.Println("Server Response:")
      fmt.Println(string(respBody))

      } else if method == "3" {
        // Get the ID of the record to update and pass it to the URL
        var param string
        fmt.Print("Enter the ID of the record to update:")
        fmt.Scanf("%v\n", &param)
        new_url := url + param
        req, err := http.NewRequest(http.MethodPut, new_url, bytes.NewBuffer(jsonData))
        if err != nil {
          fmt.Println("Error al crear la solicitud PUT:", err)
          return
        }

        // Establecer el encabezado de tipo de contenido
        req.Header.Set("Content-Type", "application/json")
        req.Header.Set("Authorization", "Bearer "+ token)

        // Realizar la solicitud PUT
        client := &http.Client{}
        resp, err := client.Do(req)
        if err != nil {
          fmt.Println("Error al realizar la solicitud PUT:", err)
          return
        }
        defer resp.Body.Close()

        // Leer la respuesta del servidor
        respBody, err := ioutil.ReadAll(resp.Body)
        if err != nil {
          fmt.Println("Error al leer la respuesta:", err)
          return
        }

        // Imprimir la respuesta del servidor
        fmt.Println("Respuesta del servidor:")
        fmt.Println(string(respBody))
      }
    } else if method == "2" {
      req, err := http.NewRequest(http.MethodGet, url, nil)
      if err != nil {
        fmt.Println("Error al crear la solicitud GET:", err)
        return
      }

      // Establecer el encabezado de autorización con el token JWT
        // var token string
        // fmt.Print("Enter your token:")
        // fmt.Scanf("%s", &token)
        req.Header.Set("Authorization", "Bearer "+ token)

      // Realizar la solicitud GET
      client := &http.Client{}
      resp, err := client.Do(req)
      if err != nil {
        fmt.Println("Error al realizar la solicitud GET:", err)
        return
      }
      defer resp.Body.Close()
      // resp, err := http.Get(url)
      // if err != nil {
      //   fmt.Println("Error:", err)
      //   return
      // }
      // defer resp.Body.Close()

      // Leer el cuerpo de la respuesta
      body, err := ioutil.ReadAll(resp.Body)
      if err != nil {
        fmt.Println("Error:", err)
        return
      }

      // Convertir el cuerpo de la respuesta a una cadena de texto
      responseBody := string(body)

      // Imprimir la respuesta
      fmt.Println(responseBody)
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

	// Realizar la solicitud DELETE
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error al realizar la solicitud DELETE:", err)
		return
	}
	defer resp.Body.Close()

	// Verificar el código de estado de la respuesta
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
