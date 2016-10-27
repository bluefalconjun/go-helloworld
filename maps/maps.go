package maps

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func Maps() {
	//map should be make before using.
	m = make(map[string]Vertex)

	//this is map insert
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	m["Google"] = Vertex{
		37.42202, -122.08408,
	}
	fmt.Println(m)

	//use map[key] to check whether map has the item.
	bellvertex := m["Bell Labs"]
	fmt.Println(bellvertex)

	//delete m[m] item.
	delete(m, "Bell Labs")

	//use double operate to check.
	bellvertexis, ok := m["Bell Labs"]
	fmt.Println(bellvertexis, ok)
	bellvertex, is := m["Bell Labs"]
	fmt.Println(bellvertex, is)
}
