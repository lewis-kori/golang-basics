package main

import "fmt"

func main() {
	// method 1
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}
	printMap(colors)

	// method 2
	// var colors map[string]string

	// method 3
	// colors := make(map[string]string)

	// colors["white"] = "#fffff"
	// fmt.Println(colors)
	// delete(colors, "white")
	// fmt.Println(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("color is " + color + " and hex is " + hex)
	}
}
