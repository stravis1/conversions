package main

import "fmt"
import "conversions/hex2base64"

func main(){
  data := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"

  fmt.Println(data)
  fmt.Println("base64 =  ", hex2base64.Hex2base64(data))
}
