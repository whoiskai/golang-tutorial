package main

import "fmt"

func main() {
	// Equivalent var init
	// var colors map[string]string
	// colors := make(map[string]string)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}

	colors["white"] = "#ffffff"
	// delete(colors, "white")

	// fmt.Println(colors)
	printMap(colors)
}

func printMap(m map[string]string) {
	for key, value := range m {
		fmt.Println("Hex code for", key, "is", value)
	}
}
