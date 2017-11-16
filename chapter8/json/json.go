package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type (
	Contact struct {
		Home  string `json:"home"`
		Phone string `json:"phone"`
	}

	Person struct {
		Name    string  `json:"name"`
		Title   string  `json:"title"`
		Contact Contact `json:"contact"`
	}
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
	var p Person

	err := json.Unmarshal([]byte(JSON), &p)
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}
	fmt.Println(p.Name)
}
