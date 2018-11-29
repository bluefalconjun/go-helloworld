/*

You're given strings J representing the types of stones that are jewels, and S representing the stones you have.  Each character in S is a type of stone you have.  You want to know how many of the stones you have are also jewels.

The letters in J are guaranteed distinct, and all characters in J and S are letters. Letters are case sensitive, so "a" is considered a different type of stone from "A".

Example 1:

Input: J = "aA", S = "aAAbbbb"
Output: 3
Example 2:

Input: J = "z", S = "ZZ"
Output: 0
Note:

S and J will consist of letters and have length at most 50.
The characters in J are distinct.

*/

import (
	"strings"
)

func numJewelsInStones(J string, S string) int {
	var nums int

	for _, OneJew := range J {
		nums += numJewInStones(OneJew, S)
	}

	return nums
}

//match one Jew only
func numJewInStones(J rune, S string) int {
	return strings.Count(S, string(J))
}