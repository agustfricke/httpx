package req

import (
	"fmt"
	"net/http"
)

func Delete(url string, token string) {
      var param string
  fmt.Print("URL: ")
      fmt.Scanf("%v\n", &param)
      new_url := url + param

      req, err := http.NewRequest(http.MethodDelete, new_url, nil)
      req.Header.Set("Authorization", "Bearer "+ token)

      if err != nil {
        fmt.Println("Error:", err)
        return
	    }

      client := &http.Client{}
      resp, err := client.Do(req)
      if err != nil {
        fmt.Println("Error:", err)
        return
      }
      defer resp.Body.Close()

      if resp.StatusCode == http.StatusOK {
        fmt.Println("La solicitud DELETE fue exitosa.")
      } else {
        fmt.Println("La solicitud DELETE response con el código de estado:", resp.StatusCode)
      }
}
