package main

import (

  "fmt"
  "unicode/utf8"
)

func strings(){

  const s = "realest"

  fmt.Println("len: ",len(s))

  for i := 0 ; i < len(s) ; i++ {
    fmt.Printf("%x ",s[i])
  }

  fmt.Println("Rune count:", utf8.RuneCountInString(s))

}
