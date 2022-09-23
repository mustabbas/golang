package main

import (
	"fmt"
	"strings"
)

func main() {
    var str string
    fmt.Println("Enter string")
    fmt.Scan(&str)
   if strings.ToUpper(string(str[0])) == "I" && strings.ToUpper(string(str[len(str)-1])) == "N" && strings.ContainsAny(strings.ToUpper(str), "A") {
    fmt.Println("Found!")
   } else{
    fmt.Println("Not Found!")
   }
}
