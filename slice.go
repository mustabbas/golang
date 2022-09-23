package main

import (
        "fmt"
	"sort"
	"strconv"
)

func main() {
var userInput string
sli:= make([]int,0,3)
i := 0
for {
   fmt.Println("please enter integers")
   fmt.Scan(&userInput)
   if userInput != "X" {
   num, err := strconv.Atoi(userInput)
   if err != nil {
       fmt.Println(err)
   }
   sli = append(sli,num)
   sort.Ints(sli)
   fmt.Println(sli)
   }else {
    break
   }
i++
}
}
