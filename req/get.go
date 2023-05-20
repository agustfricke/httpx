package req

import (
	"fmt"
	"io/ioutil"
	"net/http"

)

func Get(url string, token string) {
  var param string
  fmt.Print("URL: ")
  fmt.Scanf("%v\n", &param)
  new_url := url + param + "/"

  req, err := http.NewRequest(http.MethodGet, new_url, nil)
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


}
