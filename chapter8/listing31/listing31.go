package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	c := make(map[string]interface{})
	c["name"] = "zhangxiao"
	c["title"] = "programmer"
	c["contact"] = []map[string]interface{}{
		{
			"home": "zjj",
			"cell": "18710829722",
		},
		{
			"home": "xi'an",
			"cell": "cxxxxzx",
		},
	}
	data, err := json.Marshal(c)
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}
	fmt.Println(string(data))
}
