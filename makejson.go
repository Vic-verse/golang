package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	m := make(map[string]string)
	var x string
	var y string
	fmt.Print("Enter your name :\t")
	fmt.Scanln(&x)
	fmt.Print("Enter your address :\t")
	fmt.Scanln(&y)
	m["name"] = x
	m["address"] = y
	json, err := json.MarshalIndent(m, "", "\t")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(json))
	fmt.Println("Program closed **THANKU** and **HAVE A NICE DAY** :P")
	
}