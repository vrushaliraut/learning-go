package main

import (
	"fmt"
	"maps"
)

func mapConcept() {

	// looking up value by name rather than index - map
	// map[keyType]valueType

	// The var declaration - The Nil Map

	var nilMap map[string]int
	fmt.Print(nilMap)

	// The Map Literal
	// you can create a ready to use map by assigning it a map literal

	// An empty map literal. Not nil. Length is 0
	totalWins := map[string]int{}
	fmt.Println("Total Wins:", totalWins)

	// actual map

	teams := map[string][]string{
		"Orcas":   []string{"Fred", "Ralph", "Bijou"},
		"Lions":   []string{"Sarah", "Peter", "Billie"},
		"Kittens": []string{"Waldo", "Raul", "Ze"}, // Trailing  comma is required
	}

	fmt.Println("Teams:", teams)

	// The make function - create a map with initial capacity of 10

	ages := make(map[int][]string, 10)
	fmt.Println("ages", ages)

	// Nil vs. Empty
	var m1 map[string]int

	fmt.Println("m1 is nil :", m1 == nil)

	m2 := map[string]int{}
	fmt.Println("m2 is nil :", m2 == nil)

	literalSyntax()

	mapConstraintAndBehaviour()

	assignAndRetriveValueFromMap()

	wordFrequencyLogic()

	commaOkIdiom()

	theInventoryCheck()

	deletingFromMap()

	basicDeletion()

	lengthDistinction()

	usingMapAsASet()

	stringSet()

	fastLook()

}

/*
You have a slice of blocked IP addresses blocked := []string{"10.0.0.1", "192.168.1.1"}.
You need to check 1,000 incoming requests against this list. Why should you convert this slice to a map first?
*/
func fastLook() {
	/*
		blocked := []string{"10.0.0.1", "192.168.1.1"}
		delete(blocked, "10.0.0.1") */
}

/*
- Create a set to track unique usernames. Initialize it with "alice" and "bob". Then, add "alice" again.
- Print the length of the set to confirm duplicates are ignored.
*/
func stringSet() {
	// Initialize set with data
	users := map[string]bool{
		"alice": true,
		"bob":   true,
	}

	// Try to add duplicate
	users["alice"] = true
	fmt.Println(len(users)) // Prints 2
}

func usingMapAsASet() {
	// set guarantee one of a value - uniqueness - but doesn't guarantee any particular order
	// O(1) // set with map - map[T]bool
	// key is a type you want to store in the set - int, string // the value is bool

	// We want a set of unique integers from this list
	vals := []int{5, 10, 2, 5, 3, 4, 5, 6, 5, 3, 4, 5, 6}

	// Declare the map ( the set )
	intSet := map[int]bool{}

	// Iterate over the slice
	for _, val := range vals {
		// Assign 'true' for the key
		// if the key exists, it just overwrites 'true' with 'true'
		intSet[val] = true
	}

	//The length of vals is different and set is diff - removes duplicate

	fmt.Println(len(vals), len(intSet))

	// check for existence
	if intSet[5] {
		fmt.Println("5 is in the set")
	}

	// Check for missing value
	// Returns zero value (false) if key is missing
	if intSet[500] {
		fmt.Println("500 is in the set")
	} else {
		fmt.Println("500 is NOT in the set")
	}
}

/*
Write a program that declares a slice s := []int{1, 2} and a map m := map[int]int{1: 1, 2: 2}.
Call clear on both. Print len(s) and len(m).
*/
// Why: This reinforces the fundamental difference in how clear affects the length of these two composite types.
func lengthDistinction() {
	s := []int{1, 2}
	m := map[int]int{1: 1, 2: 2}
	m1 := map[int]int{1: 1, 2: 2}

	fmt.Println("maps equality :: ", maps.Equal(m, m1))
	clear(s) // clear on s will not clear the slice actually it will still be length of 2
	clear(m)

	fmt.Printf("Slice len: %d\n", len(s))
	fmt.Printf("Map len: %d\n", len(m))

	// nil and empty map conversion
	var nilMap map[int]int
	emptyMap := map[int]int{} // empty non-nil map

	fmt.Println(maps.Equal(nilMap, emptyMap))
}

func basicDeletion() {
	inventory := map[string]int{
		"ItemA": 10,
		"ItemB": 20,
	}
	fmt.Println("inventory basicDeletion :: ", inventory)
	delete(inventory, "ItemA")
	fmt.Println("inventory basicDeletion ::", inventory)

	counts := map[string]int{"a": 1, "b": 2}

	// Empty the map
	clear(counts)

	// We can write to it immediately because it is not nil
	counts["c"] = 3

	fmt.Println(counts)
}

func deletingFromMap() {
	// built-in delete function
	// delete takes two argument

	m := map[string]int{
		"hello": 5,
		"world": 7,
	}

	// Removes the "hello" key from the map
	delete(m, "hello")

	fmt.Println(m)

	delete(m, "any_key")
	fmt.Println("missing key ::", m)
}

/*
You are building an inventory system. inventory := map[string]int{"Apples": 5, "Bananas": 0}.

Check for "Apples".
Check for "Bananas".
Check for "Cherries". Print the value and the boolean existence flag for each.

Solution 1
Why: This reinforces that 0 is a valid value ("we have 0 Bananas in stock")
which is distinct from a missing key ("we do not carry Cherries").
*/
func theInventoryCheck() {
	inventory := map[string]int{
		"Apples":  1,
		"Bananas": 0,
	}

	// Apples
	value, ok := inventory["Apples"]
	fmt.Println("Apples:", value, ok)

	// Bananas
	value, ok = inventory["Bananas"]
	fmt.Println("Apples:", value, ok)

	// Cherry
	value, ok = inventory["Cherry"]
	fmt.Println("Cherry:", value, ok)
}

func commaOkIdiom() {
	// value - value associated with the key - or the zero value if missing
	// ok - this is bool - if key is present in the map return true else false
	/*
		value, ok := m["key"]
		fmt.Println(value, ok)  */

	fmt.Println("Distinguishing Missing Keys From Zero Values\n")
	m := map[string]int{
		"hello": 5,
		"world": 0,
	}

	// Case 1 - Key exists, non zero value

	value, ok := m["hello"]
	fmt.Println(value, ok)

	// Case 2 - Key exists, zero value
	value, ok = m["world"]
	fmt.Println(value, ok)

	// Case 2 - Key not exists, zero value
	value, ok = m["goodbye"]
	fmt.Println(value, ok)
}

/*
You are counting words. You have counts := map[string]int{}.
You encounter the word "go" for the first time. You run counts["go"] += 1.
What is the value of counts["go"] now, and why didn't this crash?
*/
func wordFrequencyLogic() {
	counts := map[string]int{}
	counts["go"] += 1
	fmt.Println("counts[\"go\"] :", counts["go"])
}

func assignAndRetriveValueFromMap() {
	// Declare and Initialize them map
	totalWins := map[string]int{}
	// write value
	totalWins["Orcas"] = 1
	totalWins["Lions"] = 2

	// read value
	fmt.Println("totalWins Orcas : ", totalWins["Orcas"])
	fmt.Println("totalWins Lions : ", totalWins["Lions"])

	/*
		Create a map named config where the key is a string (setting name) and the value is a bool (enabled/disabled).
		Initialize it with make. Set "debug" to true and "logging" to false. Print the "debug" value.
	*/

	config := make(map[string]bool, 2)
	config["debug"] = true
	config["logging"] = false

	fmt.Println("config:", config["debug"])

}

func mapConstraintAndBehaviour() {
	// Allowed - int string bool arrays pointers structs - all fields are comparable
	// Disallowed - slice, map, function
	// map themselves are not comparable

	m := map[int]int{1: 1, 2: 4, 3: 9}
	fmt.Println(len(m))

	// Reading vs writing to map - panic

	var m2 map[string]int // m2 is nil
	//read safe
	fmt.Println(m2["foo"]) // prints 0

	// write - panic - panic: assignment to entry in nil map
	// m2["bar"] = 1

	// Fixing the panic - You must initialize the map before writing
	// Option 1: Literal
	mapLiteral := map[int]int{}
	// Option 2: Make
	// m := make(map[int]int)
	mapLiteral[1] = 100 // Safe

}

/*
- Declare a map named menu where the key is a string (dish name) and the value is a float64 (price).
- Initialize it with "Soup": 10.50 and "Salad": 8.00.
*/
func literalSyntax() {
	dishName := map[string]float64{
		"Soup":  10.50,
		"Salad": 8.00,
	}

	fmt.Println(dishName)

	/*
		Why would you use make(map[string]int, 1000) instead of just map[string]int{}
		if you know you are about to add 1,000 users?

	*/
	/*Why: This is a performance optimization.
	How: Using make with a size hint allows the Go runtime to allocate enough memory for 1,000 entries upfront.
	If you use the empty literal, the map has to grow (allocate new memory, copy data, garbage collect old memory)
	multiple times as you add elements, which is slower.

	*/
}
