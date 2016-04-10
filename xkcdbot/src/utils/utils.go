package utils

import "fmt"
import "time"
import "io/ioutil"

func ParseDate(s string) time.Time {
  layout := "Mon, 2 Jan 2006 15:04:05 -0700"
  t, err := time.Parse(layout, s)

  if err != nil {
    fmt.Println("Fuck dates!")
  }

  return t
}

func ReadReferenceId() string {
  bytes, err := ioutil.ReadFile(".lastGuid")
  if err != nil {
    fmt.Println("File not exists maybe?")
    fmt.Println(err)
  }

  return string(bytes)
}

func WriteReferenceId(id string) {
  //err := ioutil.WriteFile(".lastGuid", i.Guid, 0644)
}
