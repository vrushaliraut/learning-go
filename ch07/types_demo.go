package main

import (
	"fmt"
	"strconv"
)

type score int

type Convertor func(string) score

/*type person struct {
	name string
	age  int
}*/

type TeamScore map[string]score

func convertScore(s string) score {
	if s == "" {
		return 0
	}

	value, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return score(value)

}

func main() {
	var s score = 100
	fmt.Println(s)

	var c Convertor = convertScore // functional programming
	s = c("20")
	fmt.Println(s)

	var scores TeamScore = make(map[string]score)
	scores["team1"] = 100
	scores["team2"] = 200
	fmt.Println(scores)
}
