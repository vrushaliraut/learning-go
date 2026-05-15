package main

import (
	"fmt"
	"strconv"
)

type score int

type HighScore score
type Convertor func(string) score

/*type person struct {
	name string
	age  int
}*/

func (s score) reset() {
	fmt.Println("score reset")
}

func (s HighScore) reset() {
	fmt.Println("HighScore reset")
}

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
	s.reset()
	fmt.Println(s)

	var c Convertor = convertScore // functional programming
	s = c("20")
	fmt.Println(s)

	var scores TeamScore = make(map[string]score)
	scores["team1"] = 100
	scores["team2"] = 200
	fmt.Println(scores)

	var h HighScore = 100
	h.reset()
	fmt.Println(h)

	z := HighScore(s)
	fmt.Println(z)
}
