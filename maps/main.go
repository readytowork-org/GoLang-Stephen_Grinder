package main

import "fmt"

func main() {
	// ANOTHER WAY TO DECLARE MAP
	// colors := make(map[string]string)

	// ANOTHER WAY TO DECLARE MAP
	// var colors map[string]string

	// ANOTHER WAY TO DECLARE MAP
	// interface{} accepts any data type
	colors := map[string]interface{}{
		"red":  "#0098",
		"blue": "#45gf5",
		"Code": []int{1, 2, 3},
	}

	// Add items to map
	colors["white"] = "#fffff"
	colors["green"] = "#8a799"

	// Delete value from the map
	// delete(colors, "white")

	// Iterate over the key value pair of map
	for key, value := range colors {
		fmt.Println(key, value)
	}

	fmt.Println(colors)
}
