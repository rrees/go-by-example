package main

import b64 "encoding/base64"
import "fmt"

func main() {
  data := []byte("abc123!?$*&()'-=@~")

  stringEncoding := b64.StdEncoding.EncodeToString(data)
  fmt.Println(stringEncoding)

  stringDecoding, _ := b64.StdEncoding.DecodeString(stringEncoding)
  fmt.Println(string(stringDecoding))

  urlEncoding := b64.URLEncoding.EncodeToString(data)
  fmt.Println(urlEncoding)

  urlDecoding, _ := b64.URLEncoding.DecodeString(urlEncoding)
  fmt.Println(string(urlDecoding))
}
