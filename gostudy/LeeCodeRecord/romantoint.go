package romantoint

/*
Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000

I can be placed before V (5) and X (10) to make 4 and 9.
X can be placed before L (50) and C (100) to make 40 and 90.
C can be placed before D (500) and M (1000) to make 400 and 900.
*/

import (
	"fmt"
)

func romanToInt(s string) int {
	var ret = 0

	fmt.Println("s=", s)
	const i = 0

	for {
		fmt.Println("i,s=", i, s)
		switch s[i] {
		case 'I': //special
			if i+1 < len(s) {
				if s[i+1] == 'V' {
					ret += 4
					if len(s) == 2 {
						return ret
					}
					s = s[2:]
					break
				}
				if s[i+1] == 'X' {
					ret += 9
					if len(s) == 2 {
						return ret
					}
					s = s[2:]
					break
				}

			}
			ret += 1
			if len(s) == 1 {
				return ret
			}
			s = s[1:]
		case 'V':
			ret += 5
			if len(s) == 1 {
				return ret
			}
			s = s[1:]

		case 'X': //special
			if i+1 < len(s) {
				if s[i+1] == 'L' {
					ret += 40
					if len(s) == 2 {
						return ret
					}
					s = s[2:]
					break
				}
				if s[i+1] == 'C' {
					ret += 90
					if len(s) == 2 {
						return ret
					}
					s = s[2:]
					break
				}

			}
			ret += 10
			if len(s) == 1 {
				return ret
			}
			s = s[1:]
		case 'L':
			ret += 50
			if len(s) == 1 {
				return ret
			}
			s = s[1:]
		case 'C': //special
			if i+1 < len(s) {
				if s[i+1] == 'D' {
					ret += 400
					if len(s) == 2 {
						return ret
					}
					s = s[2:]
					break
				}
				if s[i+1] == 'M' {
					ret += 900
					if len(s) == 2 {
						return ret
					}
					s = s[2:]
					break
				}

			}
			ret += 100
			if len(s) == 1 {
				return ret
			}
			s = s[1:]
		case 'D':
			ret += 500
			if len(s) == 1 {
				return ret
			}
			s = s[1:]
		case 'M':
			ret += 1000
			if len(s) == 1 {
				return ret
			}
			s = s[1:]
		default:
		}
	}
	return ret
}
