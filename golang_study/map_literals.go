package main

import "fmt"

type Veruct struct {
	Lat, Long float64
}

var mn = map[string]Veruct{
	"Bell Labs": {
		40.68433, -74.39967,
	},
	"Google": {
		37.42202, -122.08408,
	},
}

func main() {
	fmt.Println(mn)
}
