package main

import (
        "fmt"
	"os"
	"strings"
	"bufio"
)

type person struct {
  fname string
  lname string
}

func main() {
  readFile, err := os.Open("data.text")
  sli := make([]person, 0, 0)
  if err != nil {
    fmt.Println(err)
  }else {
    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)
    for fileScanner.Scan() {
        fullName := strings.Fields(string(fileScanner.Text()))
        if len(fullName) == 0 {
           break
	}else {
          p := person{fullName[0],fullName[1]}
         sli = append(sli, p)   
	}
    }
  }

  for _, v := range sli {
		fmt.Println(v.fname,v.lname)
    }
  readFile.Close()
}
