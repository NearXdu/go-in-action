package main

import "fmt"

//unsorted key-value pair
func main() {
	//key-string  value-int
	//dict := make(map[string]int)
	//dict := map[string]string{"red": "#da1337", "orange": "#e95a22"}
	//dict := map[int][]string{}

	//	colors := make(map[string]string)
	//
	//	colors["red"] = "#da1337"
	//	colors["orange"] = "#e95a22"
	//	value, exists := colors["blue"]
	//	if exists {
	//		fmt.Println(value)
	//	} else {
	//		fmt.Println("not exist!")
	//	}
	//
	//	value, exists = colors["red"]
	//	if exists {
	//		fmt.Println(value)
	//	} else {
	//		fmt.Println("not exist!")
	//	}

	//var colors map[string]string
	//colors["red"] = "#da1337"

	//fmt.Println(colors["red"])

	colors := map[string]string{
		"AliceBlue":   "#f0f8ff",
		"Coral":       "#ff7f50",
		"DarkGray":    "#a9a9a9",
		"ForestGreen": "#228b22",
	}
	for key, value := range colors {
		fmt.Printf("Key: %s \t Value: %s\n", key, value)
	}
	remove("Coral", colors)
	fmt.Println("---------------------")
	for key, value := range colors {
		fmt.Printf("Key: %s \t Value: %s\n", key, value)
	}

}

func remove(key string, m map[string]string) {
	delete(m, key)
}
