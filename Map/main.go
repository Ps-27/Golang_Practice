package main

import "fmt"

func main(){

	// var colors map[string]string
    // colors:=make(map[int]string)
	colors := map[string]string{
		"red":"#ff0000",
		"green":"#4bf745",
		"white":"#28f444",

	}

	// colors[10]="#wo97"
	// delete(colors,10)
	for i,v:=range colors{
		fmt.Printf("key %s and value %v\n", i, v)
	}
	
	fmt.Println(colors)
}