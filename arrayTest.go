package main

import "fmt"

func main(){
	planets := [...] string {
		"水星",
		"金星",
		"地球",
		"火星",
		"木星",
		"土星",
		"天王星",
		"海王星",
	}
	terrestrial := planets [0:4]
	gasGiants := planets [0:8]
	gasGiants [2] = "冥王星"
	fmt.Println(terrestrial[2],gasGiants,planets)
}