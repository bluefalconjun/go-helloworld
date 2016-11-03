package maps

import (
	"fmt"
	"strings"
)

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

func WordCount(s string) map[string]int {
	fmt.Println(s)
	var words []string = strings.Fields(s)
	retm := make(map[string]int)

	for i := 0; i < len(words); i++ {
		if _, ok := retm[words[i]]; ok {
			retm[words[i]] = retm[words[i]] + 1
			continue
		}
		retm[words[i]] = 1

	}

	return retm
}

func MapsTest() {
	fmt.Println(WordCount("I ate a donut. Then I ate another donut."))
}
