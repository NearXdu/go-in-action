package main

import (
	"encoding/json"
	"fmt"
	"log"
)

var JSON = `
{
    "name": "zhangxiao",
    "title": "programmer",
    "contact": {
        "home": "zjj",
        "phone": "18710829722"
    }
}
`

func main() {
	var c map[string]interface{}
	err := json.Unmarshal([]byte(JSON), &c)
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}
	fmt.Println("Name: \t", c["name"])
	fmt.Println("Title: \t", c["title"])

	//type assertion
	fmt.Println("Home: \t", c["contact"].(map[string]interface{})["home"])
	fmt.Println("Phone: \t", c["contact"].(map[string]interface{})["phone"])
}
