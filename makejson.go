package main

import (
	"fmt"
	"encoding/json"

)


func main() {
	var firstName string
	var address string
	var m map[string]string
	fmt.Println("Enter first name:")
	fmt.Scan(&firstName)
	fmt.Println("Enter address: ")
	fmt.Scan(&address)
	m = make(map[string]string)
	m["firstName"] = firstName
	m["address"] = address
	jsonData, err := json.Marshal(m)
	if err != nil {
	   fmt.Println(err)
	} else {
	fmt.Println(string(jsonData))
        }
}
