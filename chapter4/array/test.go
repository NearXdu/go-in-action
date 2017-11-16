package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	var array [2]string
	array2 := [...]string{"red", "green"}
	array = array2
	fmt.Println(array[0])
}
