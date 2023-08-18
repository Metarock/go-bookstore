package utils

import (
  "encoding/json"
  "io/ioutil"
  "net/http"
)

func ParseBody(request *http.Request, x interface()) {
  if body, error := ioutil.ReadAll(request.Body); err == nil{
    if error := json.Unmarshal([]byte(body), x); err != nil{
      return
    }
  }
}
